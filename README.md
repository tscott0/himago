Download high-resolution images taken by the [Himawari 8 satellite](https://en.wikipedia.org/wiki/Himawari_8). A command line tool written in Go.

Images of the Asia-Pacific region taken by Himawari 8 can be viewed online at http://himawari8.nict.go.jp/. Each image of the Earth is composed of a grid of tiles, allowing the user to freely pan and zoom. Himago downloads all images and stitches them together.

## Building

Built with Go version 1.7.3

```
$ git clone https://github.com/tscott0/himago.git 
$ cd himago
$ go build -o himago cmd/himago/main.go
```

## Usage

```
$ ./himago
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
* Bandwidth

## Known issues
* Colour
* "No Image"


## TODO
* ~~Restructure into library + command line tool~~
* Handling of "No Image" images
  * 14:40 often has these or times in the near future. Attempt rolling back 10mins at a time.
* Use the current time to nearest 10m (currently only to nearest hour)
* 404 should fail but be handled better
* Remove globals
* Support for JPEG output
* Custom output file name
* Option to save intermediate images
* Comments
* Reorder types, functions, vars
* Unit tests
  * readImage to test drawing with local files
* Debug logging mode?
* Consistent use of terminology: A tile is drawn to produce an image
* Cropping functionality
  * Just crop image before writing
  * Improve by skipping downloads for images that aren't needed.
* Create helper for URL construction (unit tested)
  * 404 error test
  * No Image test
* Measure performance
* Percentage completion in-line?
* Download speed in-line
* Summarise output image: location, file size, dimensions, format, cropping?
* --help override? If possible, add examples.
* Reorder flags
