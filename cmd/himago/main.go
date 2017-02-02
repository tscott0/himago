// The himago command-line tool
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"time"

	flag "github.com/ogier/pflag"
	"github.com/tscott0/himago"
)

var (

	// Flags

	// dateTimeFlags group all flags used to specify a date and time
	// for which to fetch images
	dateTimeFlags = flag.NewFlagSet("dateTime", flag.ExitOnError)
	now           = time.Now()

	year = dateTimeFlags.IntP("year", "y", now.Year(),
		"The year the image was taken e.g. 2016")
	month = dateTimeFlags.IntP("month", "m", int(now.Month()),
		"The month of the year the image was taken e.g. 5 means May")
	day = dateTimeFlags.IntP("day", "d", now.Day(),
		"The day of the month the image was taken e.g. 30")
	hour = dateTimeFlags.IntP("hour", "h", now.Hour(),
		"The hour the image was taken in 24-hour format e.g. 16 means 4pm")
	min = dateTimeFlags.IntP("minute", "i", now.Minute(),
		"The minute the image was taken.\n"+
			"    \tReverts to last 10min multiple e.g. 15 becomes 10")

	// outputFlags group flags which modify the output image in some way
	outputFlags = flag.NewFlagSet("output", flag.ExitOnError)

	zoom    himago.Zoom
	bandURL himago.BandURL
	outFile = outputFlags.StringP("output", "o", "output.png", "The name of the file to write to")
	red     = outputFlags.IntP("red", "R", 0, "(0-255) Red level for background color")
	green   = outputFlags.IntP("green", "G", 0, "(0-255) Green level for background color")
	blue    = outputFlags.IntP("blue", "B", 0, "(0-255) Blue level for background color")

	//cropSize  himago.Xy
	//cropStart himago.Xy
	//flag.Var(&cropSize, "cropSize",
	//"Dimensions of the cropped image in the form <width>x<height>")
	//flag.Var(&cropStart, "cropStart",
	//"Start point for cropping to cropSize in the form <xcoord>x<ycoord>")
)

func main() {
	outputFlags.VarP(&zoom, "zoom", "z", "Zoom level 1-5")
	outputFlags.VarP(&bandURL, "band", "b",
		"Electromagnetic band. Accepts integers between 1 and 16 inclusive\n"+
			"    \tIf a band is not specified a full-colour image will be produced.")

	// Override usage to be more unix-like
	flag.Usage = func() {
		usage := `usage: himago [--help] [-z zoom] [-b band] [-o output_file]
              [-y year] [-m month] [-d day] [-h hour] [-m minute] [-s second]`
		fmt.Printf("%v\n", usage)

		fmt.Printf("\n Date and time:\n")
		dateTimeFlags.PrintDefaults()

		fmt.Printf("\n Output:\n")
		outputFlags.PrintDefaults()

	}

	flag.Parse()

	// Default values for custom flags aren't supported.
	// Check whether flags are passed here and set defaults.
	if !bandURL.IsSet() {
		bandURL.Default()
	}
	if !zoom.IsSet() {
		zoom.Default()
	}

	// Set the background colour
	bg := image.NewUniform(color.RGBA{uint8(*red), uint8(*green), uint8(*blue), 255})

	// Construct a new time using the current time as the default.
	// Override with values passed on the command line.
	imageTime := himago.SatTime{
		time.Date(*year, time.Month(*month), *day, *hour, *min,
			0, 0, time.UTC),
	}

	tiles, err := himago.GetTiles(bandURL, zoom, imageTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var outImg draw.Image
	err = himago.DrawTiles(tiles, outImg, *outFile, bg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
