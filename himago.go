// Package himago provides functions to download images ultimately coming from the
// Himawari 8 satellite. Multiple smaller images, Tiles, are downloaded and stitched
// together to produce a single large image.
package himago

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"os"
)

const defaultTileSize = 550

// downloadTile will send a GET request to url and decode the response into an image
// using image.Decode.
// It returns an image.Image and any error encountered.
func downloadTile(url string) (Tile, error) {
	fmt.Printf("Downloading %v\n", url)
	var tile Tile

	response, err := http.Get(url)
	if err != nil {
		return tile, err
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}()

	newImg, _, err := image.Decode(response.Body)
	if err != nil {
		return tile, err
	}

	// Finally wrap the image.Image in a Tile and return it
	tile = Tile{newImg}
	return tile, nil
}

// Take a SatTime and construct a URL.
// Assumes that the time is valid.
func urlFromSatTime(band Band, t SatTime, gridWidth, i, j int) string {
	return fmt.Sprintf(band.URL(),
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
func GetTiles(band Band, zoom Zoom, imageTime SatTime) ([][]Tile, error) {
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

			url := urlFromSatTime(band, imageTime, gridWidth, i, j)
			tile, err := downloadTile(url)

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
						url = urlFromSatTime(band, imageTime, gridWidth, i, j)
						tile, err = downloadTile(url)

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
func DrawTiles(band Band, tiles [][]Tile, outImg draw.Image, fileName string, bg Color, fg Color) error {
	// Set the background colour
	backdrop := image.NewUniform(bg)

	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// Assume images are always square
	gridWidth := len(tiles)

	// Create a new image with a black background
	imgRect := image.Rect(0, 0, gridWidth*defaultTileSize, gridWidth*defaultTileSize)
	outImg = image.NewRGBA(imgRect)

	draw.Draw(outImg, outImg.Bounds(), backdrop, image.ZP, draw.Src)

	// Loop over the Tiles and Draw them
	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridWidth; y++ {
			// Define the bounds of the image.Rectangle for this Tile
			tileRect := image.Rect(
				x*defaultTileSize,
				y*defaultTileSize,
				(x+1)*defaultTileSize,
				(y+1)*defaultTileSize)

			// Full colour images have no transparency
			// Only set the foreground colour when using a band
			if band != Band(0) {
				tiles[x][y].setForeground(fg)
			}

			// Draw the Tile to the Image
			draw.Draw(outImg,
				tileRect,
				image.Image(tiles[x][y]),
				image.ZP,
				draw.Over)
		}
	}

	// Write the image to file
	//err = jpeg.Encode(outFile, outImg, &jpeg.Options{90})
	err = png.Encode(outFile, outImg)
	if err != nil {
		return err
	}

	fmt.Printf("\nSaved to %v\n", fileName)

	return nil
}
