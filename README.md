# himago

[![GoDoc](https://godoc.org/github.com/tscott0/himago?status.svg)](https://godoc.org/github.com/tscott0/himago) [![Go Report Card](https://goreportcard.com/badge/github.com/tscott0/himago)](https://goreportcard.com/report/github.com/tscott0/himago) [![Build Status](https://travis-ci.org/tscott0/himago.svg?branch=master)](https://travis-ci.org/tscott0/himago)
---
Download high-resolution images taken by the [Himawari 8 satellite](https://en.wikipedia.org/wiki/Himawari_8). A command line tool written in Go.

## Overview 

Images of the Asia-Pacific region taken by Himawari 8 can be viewed online at http://himawari8.nict.go.jp/. Each image of the Earth is composed of a grid of Tiles, allowing the user to freely pan and zoom. Himago downloads all images and stitches them together.

## Examples
<img src="/_examples/arch.png?raw=true" width="25%"><img src="/_examples/b05-06.png?raw=true" width="25%"><img src="/_examples/b11-12.png?raw=true" width="25%"><img src="/_examples/blue.png?raw=true" width="25%">
<img src="/_examples/current.png?raw=true" width="25%"><img src="/_examples/dominos.png?raw=true" width="25%"><img src="/_examples/facebook.png?raw=true" width="25%"><img src="/_examples/flickr.png?raw=true" width="25%">
<img src="/_examples/gplus.png?raw=true" width="25%"><img src="/_examples/heineken.png?raw=true" width="25%"><img src="/_examples/ikea.png?raw=true" width="25%"><img src="/_examples/lego.png?raw=true" width="25%">
<img src="/_examples/noband-12.png?raw=true" width="25%"><img src="/_examples/orange.png?raw=true" width="25%"><img src="/_examples/purple.png?raw=true" width="25%"><img src="/_examples/python.png?raw=true" width="25%">
<img src="/_examples/red.png?raw=true" width="25%"><img src="/_examples/reddit.png?raw=true" width="25%"><img src="/_examples/rgb1.png?raw=true" width="25%">

## Install

```
go get github.com/tscott0/himago
```

## Build

```
$ git clone https://github.com/tscott0/himago.git
$ cd himago/cmd/himago/
$ go build
```

## Usage

```
usage: himago [--help] [-z zoom] [-b band] [-o output_file]
              [-y year] [-m month] [-d day] [-h hour] [-m minute] [-s second]

  -b, --band value
    	Electromagnetic band. Accepts integers between 1 and 16 inclusive
    	If a band is not specified a full-colour image will be produced.
  -B, --blue int
    	(0-255) Blue level for background color
  -d, --day int
    	The day of the month the image was taken e.g. 30 (default 3)
  -G, --green int
    	(0-255) Green level for background color
  -h, --hour int
    	The hour the image was taken in 24-hour format e.g. 16 means 4pm (default 19)
  -i, --minute int
    	The minute the image was taken.
    	Reverts to last 10min multiple e.g. 15 becomes 10 (default 22)
  -m, --month int
    	The month of the year the image was taken e.g. 5 means May (default 2)
  -o, --output string
    	The name of the file to write to (default "output.png")
  -R, --red int
    	(0-255) Red level for background color
  -y, --year int
    	The year the image was taken e.g. 2016 (default 2017)
  -z, --zoom value
    	Zoom level 1-5 (default )
```

### Zoom
```
Zoom  Grid   Resolution
1     1x1    550  x 550
2     2x2    1100 x 1100
3     4x4    2200 x 2200
4     8x8    4400 x 4400
5     16x16  8800 x 8800
```
*Default zoom is 2*

### Considerations
* Bandwidth: The [JMA](https://en.wikipedia.org/wiki/Japan_Meteorological_Agency) have made generously made these images freely available. While this tool might be useful for wallpapers, please don't abuse it by downloading hi-resolution images regularly. Please use responsinbly.


## Acknowledgements
* [Japan Meteorological Agency](https://en.wikipedia.org/wiki/Japan_Meteorological_Agency)
* [NICT](https://www.nict.go.jp/en/about/)
* [Michael Pote](https://github.com/MichaelPote) created a [script](https://gist.github.com/MichaelPote/92fa6e65eacf26219022) that inspired many similar tools like this.
* [Jacob Kelley](https://github.com/jakiestfu) for his excellent work on [himawari.js](https://github.com/jakiestfu/himawari.js)

## Known issues
* Unrealistic colours: According to [Wikipedia](https://en.wikipedia.org/wiki/Himawari_8), the images returned are true-colour. Looking at the colour of Australia, in particular, the colours don't look accurate. Correcting the colour to make it appear more natural looks complicated.
* Occasionally will get 404 errors. Himago doesn't handle these automatically so it would require the user to specify a different date or time.

## TODO
* Pass number of rollback attempts on the command line. Maximum?
* JPEG output format (inferred from filename -o)
* Percentage completion in-line?
* Download speed in-line
* Summarise output image: location, file size, dimensions, format, cropping?
* Cropping functionality
  * Just crop image before writing
* Option to save intermediate Tile images
  * Improve by skipping downloads for images that aren't needed.
* --version
* 404 should fail but be handled better
* Consider using https://github.com/pkg/errors
* Measure performance


## License

MIT.
