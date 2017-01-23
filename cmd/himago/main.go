// The himago command-line tool
package main

import (
	"flag"
	"fmt"
	"image/draw"
	"os"
	"time"

	himago "../../../himago"
)

var (
	cropSize  himago.Xy
	cropStart himago.Xy
	imageTime himago.SatTime
	zoom      int
	outImg    draw.Image
)

func init() {
	now := time.Now()

	flag.Var(&cropSize, "cropSize",
		"Dimensions of the cropped image in the form <width>x<height>")
	flag.Var(&cropStart, "cropStart",
		"Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
	flag.IntVar(&zoom, "zoom", 2,
		"Zoom level 1-5")

	year := flag.Int("year", now.Year(), "Year of the image.")
	month := flag.Int("month", int(now.Month()), "Month of the image.")
	day := flag.Int("day", now.Day(), "Day of the image.")
	hour := flag.Int("hour", now.Hour(), "Hour of the image.")
	min := flag.Int("minute", now.Minute(), "Minute of the image.")
	// Construct a new time using the current time as the default.
	// Override with values passed on the command line.
	imageTime = himago.SatTime{
		time.Date(*year, time.Month(*month), *day, *hour, *min,
			0, 0, time.UTC),
	}
}

func main() {

	flag.Parse()

	tiles, err := himago.GetTiles(zoom, imageTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = himago.DrawTiles(tiles, outImg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
