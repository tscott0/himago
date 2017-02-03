<img src="http://i.imgur.com/XBHi48b.png" width="100%">
[![GoDoc](https://godoc.org/github.com/tscott0/himago?status.svg)](https://godoc.org/github.com/tscott0/himago) [![Go Report Card](https://goreportcard.com/badge/github.com/tscott0/himago)](https://goreportcard.com/report/github.com/tscott0/himago) [![Build Status](https://travis-ci.org/tscott0/himago.svg?branch=master)](https://travis-ci.org/tscott0/himago)
---
Download high-resolution images taken by the [Himawari 8 satellite](https://en.wikipedia.org/wiki/Himawari_8). A command line tool written in Go.

## Overview 

Images of the Asia-Pacific region taken by Himawari 8 can be viewed online at http://himawari8.nict.go.jp/. Each image of the Earth is composed of a grid of Tiles, allowing the user to freely pan and zoom. Himago downloads all images and stitches them together.

## Install

```
go get github.com/tscott0/himago
```

## Build

Built with Go version 1.7.3

```
$ git clone https://github.com/tscott0/himago.git
$ cd himago
$ go build -o himago cmd/himago/main.go
```

## Usage

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

## Examples
With no arguments himago will get the most recent images with a default zoom of 2.
```
$ himago
```
<img src="http://i.imgur.com/trvX2su.png" width="50%">
---
Full colour image from 12PM
```
$ himago --zoom 1 -hour 12 --minute 29 --band 11
```
<img src="http://i.imgur.com/mEeBerP.png" width="50%">
---
Full colour image from 6AM
```
$ himago --zoom 1 -hour 6
```
<img src="http://i.imgur.com/FiSLobt.png" width="50%">
---
Band 11 at 12PM

```
$ himago --zoom 1 -hour 12
```
<img src="http://i.imgur.com/I2sppS2.png" width="50%">
---
Band 5 at 6AM
```
$ himago --zoom 1 -hour 06 --band 05
```
<img src="http://i.imgur.com/ZUDIktb.png" width="50%">
---


## Acknowledgements
* [Japan Meteorological Agency](https://en.wikipedia.org/wiki/Japan_Meteorological_Agency)
* [NICT](https://www.nict.go.jp/en/about/)
* [Michael Pote](https://github.com/MichaelPote) created a [script](https://gist.github.com/MichaelPote/92fa6e65eacf26219022) that inspired many similar tools like this.
* [Jacob Kelley](https://github.com/jakiestfu) for his excellent work on [himawari.js](https://github.com/jakiestfu/himawari.js)

## Known issues
* Unrealistic colours: According to [Wikipedia](https://en.wikipedia.org/wiki/Himawari_8), the images returned are true-colour. Looking at the colour of Australia, in particular, the colours don't look accurate. Correcting the colour to make it appear more natural looks complicated.
* Occasionally will get 404 errors. Himago doesn't handle these automatically so it would require the user to specify a different date or time.

## TODO

### New features
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

### Bugs/improvements/other
* Add colour examples to README
* 404 should fail but be handled better
* Unit tests
* Debug logging
* Consider using https://github.com/pkg/errors
* Measure performance

## License

MIT.
