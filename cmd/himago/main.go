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

	now   = time.Now()
	year  = flag.Int("year", now.Year(), "Year of the image")
	month = flag.Int("month", int(now.Month()), "Month of the image")
	day   = flag.Int("day", now.Day(), "Day of the image")
	hour  = flag.Int("hour", now.Hour(), "Hour of the image")
	min   = flag.Int("minute", now.Minute(), "Minute of the image")

	zoom = flag.Int("zoom", 2, "Zoom level 1-5")

	//cropSize  himago.Xy
	//cropStart himago.Xy
	//flag.Var(&cropSize, "cropSize",
	//"Dimensions of the cropped image in the form <width>x<height>")
	//flag.Var(&cropStart, "cropStart",
	//"Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
)

func main() {
	flag.Parse()

	// Construct a new time using the current time as the default.
	// Override with values passed on the command line.
	imageTime = himago.SatTime{
		time.Date(*year, time.Month(*month), *day, *hour, *min,
			0, 0, time.UTC),
	}

	tiles, err := himago.GetTiles(*zoom, imageTime)
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
