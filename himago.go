// Package himago provides functions to download images ultimately coming from the
// Himawari 8 satellite. Multiple smaller images, referred to as tiles, are stitched
// together to produce a single large image.
package himago

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
)

// GetTile will send a GET request to url and decode the response into an image
// using image.Decode.
// It returns an image.Image and any error encountered.
func GetTile(url string) (Tile, error) {
	fmt.Printf("Downloading %v\n", url)
	var tile Tile

	response, err := http.Get(url)
	if err != nil {
		return tile, errors.New("Failed to get image from: " + url)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}()

	// Extract the response body so we can hash it
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tile, errors.New("Failed to read image body")
	}

	// Store the hex md5sum
	md5 := fmt.Sprintf("%x", md5.Sum(body))

	//newImg, _, err := image.Decode(response.Body)
	newImg, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		return tile, errors.New("Failed to read decode image from url: " + url)
	}

	// Finally wrap the image.Image in a Tile and return it
	tile = Tile{newImg, md5}
	return tile, nil
}

// Take a SatTime and construct a URL.
// Assumes that the time is valid.
func urlFromSatTime(url BandURL, t SatTime, gridWidth, i, j int) string {
	return fmt.Sprintf(string(url),
		gridWidth,
		t.Year(),
		int(t.Month()),
		t.Day(),
		t.Hour(),
		t.Minute(),
		j,
		i)
}

// GetTiles retrieves the individual tiles to construct an image at the
// required zoom level.
// Tiles are always the same size: 550x550 pixels
// A higher zoom will require more tiles to be downloaded and will
// produce a higher resolution image.
//
// Zoom  Grid   Resolution
// 1     1x1    550  x 550
// 2     2x2    1100 x 1100
// 3     4x4    2200 x 2200
// 4     8x8    4400 x 4400
// 5     16x16  8800 x 8800
func GetTiles(bURL BandURL, zoom Zoom, imageTime SatTime) ([][]Tile, error) {
	gridWidth := zoom.GridWidth()

	tiles := [][]Tile{}

	// Round down to the nearest 10 minutes
	imageTime.Round()

	// On attempting to download the first tile for an image,
	// if a "No Image" is detected then roll back 10 minutes
	// and try again. Try 3 times and then error.
	firstTile := true
	remainingRollbacks := 3

	for j := 0; j < gridWidth; j++ {
		row := []Tile{}
		for i := 0; i < gridWidth; i++ {

			url := urlFromSatTime(bURL, imageTime, gridWidth, i, j)
			tile, err := GetTile(url)

			if err != nil {
				return tiles, err
			}

			// Only perform rollback check on the first tile.
			// Assumes all tiles to be "No Image" if the first one is.
			if firstTile {
				for remainingRollbacks > 0 {
					if tile.IsNoImage() {
						fmt.Println("Bad image, rolling back.")
						imageTime.Rollback()

						// Regenerate the URL will the new time
						url = urlFromSatTime(bURL, imageTime, gridWidth, i, j)
						tile, err = GetTile(url)

						if err != nil {
							return tiles, err
						}
					}
					remainingRollbacks--
				}
			}

			// Add the tile to the array
			row = append(row, tile)
			firstTile = false
		}
		tiles = append(tiles, row)
	}

	return tiles, nil

}

// DrawTiles takes a collection of Tiles and writes them to file.
func DrawTiles(tiles [][]Tile, outImg draw.Image, fileName string) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	w := 550

	// Assume images are always square
	gridWidth := len(tiles)

	// Create a new image with a black background
	outImg = image.NewRGBA(image.Rect(0, 0, gridWidth*w, gridWidth*w))
	draw.Draw(outImg, outImg.Bounds(), image.Black, image.ZP, draw.Src)

	// Loop over the Tiles and Draw them
	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridWidth; y++ {
			draw.Draw(outImg, image.Rect(x*w, y*w, (x+1)*w, (y+1)*w), image.Image(tiles[x][y]), image.ZP, draw.Over)
		}
	}

	// Write the image to file
	err = png.Encode(outFile, outImg)
	if err != nil {
		return err
	}

	fmt.Println("Saved to output.png")

	return nil
}
