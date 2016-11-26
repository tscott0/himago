// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

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

type xy struct {
	X int
	Y int
}

func (coord *xy) String() string {
	return fmt.Sprintf("%vx%v", coord.X, coord.Y)
}

func (coord *xy) Set(value string) error {
	value = strings.ToLower(value)
	coords := strings.Split(value, "x")

	if len(coords) != 2 {
		return errors.New("Invalid coordinates.")
	}

	var err error
	coord.X, err = strconv.Atoi(coords[0])
	if err != nil {
		return errors.New("X was an invalid number. Received: " + value)
	}

	coord.Y, err = strconv.Atoi(coords[1])
	if err != nil {
		return errors.New("Y was an invalid number. Received: " + value)
	}

	if coord.X < 0 || coord.Y < 0 {
		return errors.New("Coordinates cannot be negative.")
	}

	return nil
}

type satTime struct {
	year   *int
	month  *int
	day    *int
	hour   *int
	minute *int
}

var (
	cropSize  xy
	cropStart xy
	imageTime satTime
	zoom      int
	outImg    draw.Image
)

func getTiles() ([][]image.Image, error) {
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

func drawTiles(tiles [][]image.Image) error {
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

func init() {
	now := time.Now()

	flag.Var(&cropSize, "cropSize", "Dimensions of the cropped image in the form <width>x<height>")
	flag.Var(&cropStart, "cropStart", "Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
	flag.IntVar(&zoom, "zoom", 2, "Zoom factor 1-5")
	imageTime.year = flag.Int("year", now.Year(), "Year of the image.")
	imageTime.month = flag.Int("month", int(now.Month()), "Month of the image.")
	imageTime.day = flag.Int("day", now.Day(), "Day of the image.")
	imageTime.hour = flag.Int("hour", now.Hour(), "Hour of the image.")
	imageTime.minute = flag.Int("minute", now.Minute(), "Minute of the image.")

}

func main() {

	flag.Parse()

	tiles, err := getTiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = drawTiles(tiles)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
