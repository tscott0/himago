// The himago command-line tool
package main

import (
	"fmt"
	"image/draw"
	"os"
	"time"

	flag "github.com/ogier/pflag"
	"github.com/tscott0/himago"
)

var (

	// Flags
	now = time.Now()

	year = flag.IntP("year", "y", now.Year(),
		"The year the image was taken e.g. 2016")
	month = flag.IntP("month", "m", int(now.Month()),
		"The month of the year the image was taken e.g. 5 means May")
	day = flag.IntP("day", "d", now.Day(),
		"The day of the month the image was taken e.g. 30")
	hour = flag.IntP("hour", "h", now.Hour(),
		"The hour the image was taken in 24-hour format e.g. 16 means 4pm")
	min = flag.IntP("minute", "i", now.Minute(),
		"The minute the image was taken.\n"+
			"    \tReverts to last 10min multiple e.g. 15 becomes 10")

	zoom    himago.Zoom
	band    himago.Band
	bg      himago.Color
	fg      himago.Color
	outFile = flag.StringP("output", "o", "output.png", "The name of the file to write to")

	//cropSize  himago.Xy
	//cropStart himago.Xy
	//flag.Var(&cropSize, "cropSize",
	//"Dimensions of the cropped image in the form <width>x<height>")
	//flag.Var(&cropStart, "cropStart",
	//"Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
)

func main() {
	// Set default values for custom flags
	fg.Set("#FFFFFF")
	zoom.Set("2")

	flag.VarP(&zoom, "zoom", "z", "Zoom level 1-5")
	flag.VarP(&band, "band", "b",
		"Electromagnetic band. Accepts integers between 1 and 16 inclusive\n"+
			"    \tIf a band is not specified a full-colour image will be produced.")

	flag.VarP(&bg, "bg", "B", "The background colour in hex format")
	flag.VarP(&fg, "fg", "F", "The foreground colour in hex format")
	// Override usage to be more unix-like
	flag.Usage = func() {
		usage := `usage: himago [--help] [-z zoom] [-b band] [-o output_file]
              [-y year] [-m month] [-d day] [-h hour] [-m minute] [-s second]`
		fmt.Printf("%v\n\n", usage)

		flag.PrintDefaults()
	}

	flag.Parse()

	// Construct a new time using the current time as the default.
	// Override with values passed on the command line.
	imageTime := himago.SatTime{
		time.Date(*year, time.Month(*month), *day, *hour, *min,
			0, 0, time.UTC),
	}

	tiles, err := himago.GetTiles(band, zoom, imageTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var outImg draw.Image
	err = himago.DrawTiles(band, tiles, outImg, *outFile, bg, fg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
