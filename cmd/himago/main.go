// The himago command-line tool
package main

import (
	"flag"
	"fmt"
	"image/draw"
	"os"
	"time"

	"github.com/tscott0/himago"
)

var (
	imageTime himago.SatTime
	outImg    draw.Image

	// Flags
	zoom    himago.Zoom
	bandURL himago.BandURL

	now  = time.Now()
	year = flag.Int("year", now.Year(),
		"The year the image was taken e.g. 2016")
	month = flag.Int("month", int(now.Month()),
		"The month of the year the image was taken e.g. 5 means May")
	day = flag.Int("day", now.Day(),
		"The day of the month the image was taken e.g. 30")
	hour = flag.Int("hour", now.Hour(),
		"The hour the image was taken in 24-hour format e.g. 16 means 4pm")
	min = flag.Int("minute", now.Minute(),
		"The minute the image was taken.\n"+
			"    \tReverts to last 10min multiple e.g. 15 becomes 10")

	//zoom = flag.Int("zoom", 2, "Zoom level 1-5")

	//cropSize  himago.Xy
	//cropStart himago.Xy
	//flag.Var(&cropSize, "cropSize",
	//"Dimensions of the cropped image in the form <width>x<height>")
	//flag.Var(&cropStart, "cropStart",
	//"Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
)

func main() {
	flag.Var(&zoom, "zoom", "Zoom level 1-5")
	flag.Var(&bandURL, "band",
		"Electromagnetic `band`. Accepts values 01,02...16 or \"standard\"\n"+
			"    \tNumbers must be zero-padded (default \"standard\")")

	flag.Parse()

	// Default values for custom flags aren't supported.
	// Check whether flags are passed here and set defaults.
	if !bandURL.IsSet() {
		bandURL.Default()
	}

	if !zoom.IsSet() {
		zoom.Default()
	}

	// Construct a new time using the current time as the default.
	// Override with values passed on the command line.
	imageTime = himago.SatTime{
		time.Date(*year, time.Month(*month), *day, *hour, *min,
			0, 0, time.UTC),
	}

	tiles, err := himago.GetTiles(bandURL, zoom, imageTime)
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
