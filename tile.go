package himago

import "image"

// md5sum of a known bad image ("No Image")
const noImageMD5 string = "b697574875d3b8eb5dd80e9b2bc9c749"

// Tile wraps an image.Image and provides helper functions to detect
// "no image" images.
//
// Tiles are always the same size: 550x550 pixels
type Tile struct {
	image.Image
	md5 string // The hex representation of the md5sum
}

// IsNoImage returns true if the md5sum of the image matches
// the no image hash
func (t *Tile) IsNoImage() bool {
	return t.md5 == noImageMD5
}
