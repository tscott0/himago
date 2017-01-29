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
* Bandwidth: The [JMA](https://en.wikipedia.org/wiki/Japan_Meteorological_Agency) have made generously made these images freely available. While this tool might be useful for wallpapers, please don't abuse it by downloading hi-resolution images regularly. Please use responsinbly

## Examples
<img src="http://i.imgur.com/w8dfDX5.jpg" width="50%"><img src="http://i.imgur.com/G5dK3YD.png" width="50%">

## Acknowledgements
* [Japan Meteorological Agency](https://en.wikipedia.org/wiki/Japan_Meteorological_Agency)
* [NICT](https://www.nict.go.jp/en/about/)
* [Michael Pote](https://github.com/MichaelPote) created a [script](https://gist.github.com/MichaelPote/92fa6e65eacf26219022) that inspired many similar tools like this.
* [Jacob Kelley](https://github.com/jakiestfu) for his excellent work on [himawari.js](https://github.com/jakiestfu/himawari.js)

## Known issues
* Unrealistic colours: According to [Wikipedia](https://en.wikipedia.org/wiki/Himawari_8), the images returned are true-colour. Looking at the colour of Australia, in particular, the colours don't look accurate. Correcting the colour to make it appear more natural looks complicated.

## TODO
* Handling of "No Image" images
  * ~~Could check hash of image~~ Now checks md5 for "No Image"
  * ~~Attempt rolling back 10mins at a time.~~
  * Tune rollback attempts. Currently set to 3.
  * Pass number of attempts on the command line.
* Make Zoom a type. Validate input with Flag interface.
* Add more specific examples including for bands. Also include command used to generate it.
* Ability to specify background colour when using specifying a band.
* 404 should fail but be handled better
* Support for JPEG output
* Custom output file name
* Option to save intermediate Tile images
* Unit tests
* Debug logging
* Consider using https://github.com/pkg/errors
* Cropping functionality
  * Just crop image before writing
  * Improve by skipping downloads for images that aren't needed.
* Measure performance
* Percentage completion in-line?
* Download speed in-line
* Summarise output image: location, file size, dimensions, format, cropping?
* Support for getting only a specific light band
