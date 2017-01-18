package himago

import (
	"fmt"
	"image"
)

const noImageMD5 string = "b697574875d3b8eb5dd80e9b2bc9c749"

// Tile wraps an image.Image and provides helper functions to detect
// "no image" images.
type Tile struct {
	image.Image
	md5 string // The hex representation of the md5sum
}

// Returns true if the md5sum of the image matches the no image hash
func (t *Tile) IsNoImage() bool {

	fmt.Println("MD5: " + t.md5 + "\nvs.  " + noImageMD5)
	return t.md5 == noImageMD5
}
