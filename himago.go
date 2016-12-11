// Package himago provides functions to download images ultimately coming from the
// Himawari 8 satellite. Multiple smaller images, referred to as tiles, are stitched
// together to produce a single large image.
package himago

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"net/http"
	"os"
)

type Xy struct {
	X int
	Y int
}

// getTile will send a GET request to url and decode the response into an image
// using image.Decode.
// It returns an image.Image and any error encountered.
func getTile(url string) (image.Image, error) {
	fmt.Printf("Downloading %v\n", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Failed to get image from: " + url)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}()

	newImg, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, errors.New("Failed to read decode image from url: " + url)
	}

	return newImg, nil
}

type satTime struct {
	year   *int
	month  *int
	day    *int
	hour   *int
	minute *int
}

func getTiles(zoom int, imageTime satTime) ([][]image.Image, error) {
	// Zoom level   Grid
	// 1            1x1
	// 2            2x2
	// 3            4x4
	// 4            8x8
	// 5            16x16
	gridWidth := int(math.Pow(2, float64(zoom-1)))

	tiles := [][]image.Image{}

	for j := 0; j < gridWidth; j++ {
		row := []image.Image{}
		for i := 0; i < gridWidth; i++ {
			url := fmt.Sprintf("http://himawari8-dl.nict.go.jp/himawari8/img/D531106/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png",
				gridWidth,
				*imageTime.year,
				*imageTime.month,
				*imageTime.day,
				03,
				00,
				j,
				i)

			tile, err := getTile(url)
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

func drawTiles(tiles [][]image.Image, outImg draw.Image) error {
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
			draw.Draw(outImg, image.Rect(x*w, y*w, (x+1)*w, (y+1)*w), tiles[x][y], image.ZP, draw.Src)
		}
	}

	err = png.Encode(outFile, outImg)
	if err != nil {
		return err
	}

	fmt.Println("Generated image to output.png")

	return nil
}
