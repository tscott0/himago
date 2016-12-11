// The himago command-line tool
package main

import (
	"errors"
	"flag"
	"fmt"
	"image/draw"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tscott0/himago"
)

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

var (
	cropSize  xy
	cropStart xy
	imageTime satTime
	zoom      int
	outImg    draw.Image
)

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

	tiles, err := himago.getTiles(zoom, imageTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = himago.drawTiles(tiles, outImg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
