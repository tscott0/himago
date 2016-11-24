// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"image/draw"
	"image/png"
	// _ "image/jpeg"
)

//func stitch(x int, y int, images []image.Image) {
//for _, img := range images {

//}
//}

func readImage(img string) image.Image {
	reader, err := os.Open(img)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	fmt.Printf("%v: %vx%v\n", img, bounds.Max.X, bounds.Max.Y)

	return m
}

func getImage(url string) (image.Image, error) {
	fmt.Printf("Downloading %v\n", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Failed to get image from: " + url)
	}

	defer response.Body.Close()

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
	outImg    image.Image
)

func getTiles() ([]image.Image, error) {

	xTiles, yTiles := 2 ^ (zoom - 1)

	// Make an array of images big enough to store all the tiles
	tiles := make([]image.Image, xTiles*yTiles, xTiles*yTiles)
	tileIdx := 0

	for j := 0; j < yTiles; j++ {
		for i := 0; i < xTiles; i++ {
			// TODO: fix zero-padding on numbers with fmt.Sprintf
			url := fmt.Sprintf("URL = http://himawari8-dl.nict.go.jp/himawari8/img/D531106/2d/550/%v/%v/%v/%v%v00_%v_%v.png",
				*imageTime.year,
				*imageTime.month,
				*imageTime.day,
				19,
				00,
				i,
				j)

			fmt.Println(url)

			tile, err := getImage(url)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Add the tile to the array
			tiles[tileIdx] = tile
		}
	}

}

func init() {
	now := time.Now()

	flag.Var(&cropSize, "cropSize", "Dimensions of the cropped image in the form <width>x<height>")
	flag.Var(&cropStart, "cropStart", "Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
	flag.IntVar(&zoom, 2, "Zoom factor 1-5")
	imageTime.year = flag.Int("year", now.Year(), "Year of the image.")
	imageTime.month = flag.Int("month", int(now.Month()), "Month of the image.")
	imageTime.day = flag.Int("day", now.Day(), "Day of the image.")
	imageTime.hour = flag.Int("hour", now.Hour(), "Hour of the image.")
	imageTime.minute = flag.Int("minute", now.Minute(), "Minute of the image.")

}

func main() {

	flag.Parse()

	outFile, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outImg = image.NewRGBA(image.Rect(0, 0, 1100, 550))
	draw.Draw(outImg, outImg.Bounds(), image.Black, image.ZP, draw.Src)

	//m := readImage("221000_0_0.png")
	//draw.Draw(outImg, image.Rect(0, 0, 550, 550), m, image.ZP, draw.Src)

	//b := readImage("221000_1_0.png")
	draw.Draw(outImg, image.Rect(550, 0, 1100, 550), b, image.ZP, draw.Src)

	err = png.Encode(outFile, outImg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generated image to output.png \n")
}
