package himago

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"image/png"
)

// md5sum of a known bad image ("No Image")
const noImageMD5 string = "b5fd2ee42ee01da39dbd477e9fe981cb"

// Tile wraps an image.Image and provides helper functions to detect
// "no image" images.
//
// Tiles are always the same size: 550x550 pixels
type Tile struct {
	image.Image
}

// IsNoImage returns true if the md5sum of the image matches
// the "No Image" hash.
func (t *Tile) IsNoImage() bool {
	md5sum := t.md5Sum()

	return md5sum == noImageMD5
}

// md5Sum re-encodes the image as a PNG and returns the md5sum
// of an image in hex format.
func (t *Tile) md5Sum() string {
	var b bytes.Buffer
	png.Encode(&b, t.Image)

	// Convert to hex for comparison
	return fmt.Sprintf("%x", md5.Sum(b.Bytes()))
}
