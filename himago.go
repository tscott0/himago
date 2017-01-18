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
	"math"
	"net/http"
	"os"
)

// getTile will send a GET request to url and decode the response into an image
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

func GetTiles(zoom int, imageTime SatTime) ([][]Tile, error) {
	// Zoom level   Grid
	// 1            1x1
	// 2            2x2
	// 3            4x4
	// 4            8x8
	// 5            16x16
	gridWidth := int(math.Pow(2, float64(zoom-1)))

	urlString := "http://himawari8-dl.nict.go.jp/himawari8/img/D531106/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"

	tiles := [][]Tile{}

	imageTime.Round()

	for j := 0; j < gridWidth; j++ {
		row := []Tile{}
		for i := 0; i < gridWidth; i++ {
			url := fmt.Sprintf(urlString,
				gridWidth,
				*imageTime.Year,
				*imageTime.Month,
				*imageTime.Day,
				*imageTime.Hour,
				//00,
				*imageTime.Minute,
				j,
				i)

			tile, err := GetTile(url)

			if tile.IsNoImage() {
				fmt.Println("OMG, I've found one!")
			}

			if err != nil {
				return tiles, err
			}

			// Add the tile to the array
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	return tiles, nil

}

func DrawTiles(tiles [][]Tile, outImg draw.Image) error {
	outFile, err := os.Create("./output.png")
	if err != nil {
		return err
	}

	w := 550

	// Assume images are always square
	gridWidth := len(tiles)

	outImg = image.NewRGBA(image.Rect(0, 0, gridWidth*w, gridWidth*w))

	// Black background
	draw.Draw(outImg, outImg.Bounds(), image.White, image.ZP, draw.Src)

	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridWidth; y++ {
			draw.Draw(outImg, image.Rect(x*w, y*w, (x+1)*w, (y+1)*w), image.Image(tiles[x][y]), image.ZP, draw.Src)
		}
	}

	err = png.Encode(outFile, outImg)
	if err != nil {
		return err
	}

	fmt.Println("Saved to output.png")

	return nil
}
