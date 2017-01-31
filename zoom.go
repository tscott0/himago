package himago

import (
	"errors"
	"math"
	"strconv"
)

// Zoom is an int restricted to numbers 1-5 inclusive
type Zoom int

// String returns Zoom as a string
func (z *Zoom) String() string {
	return string(*z)
}

// Set will check the value of the zoom and error if invalid
// Zoom can only be 1-5 inclusive.
func (z *Zoom) Set(zoomString string) error {
	// Attempt to cast to int
	i, err := strconv.Atoi(zoomString)
	*z = Zoom(i)

	// If it's not an integer or isn't between 1 and 5 (inclusive) error
	if err != nil || *z < 1 || *z > 5 {
		return errors.New("zoom must be an integer between 1 and 5")
	}

	return nil
}

// GridWidth returns the number of Tiles the image is square.
//
// Zoom  Grid   Resolution
// 1     1x1    550  x 550
// 2     2x2    1100 x 1100
// 3     4x4    2200 x 2200
// 4     8x8    4400 x 4400
// 5     16x16  8800 x 8800
func (z *Zoom) GridWidth() int {
	return int(math.Pow(2, float64(*z-1)))
}

// IsSet returns true if the Zoom is the non-default value.
// If the underlying int is 0 then return false.
func (z *Zoom) IsSet() bool {
	return int(*z) != 0
}

// Default sets the BandURL to the default value
// Calling Default is required IsSet returns false
func (z *Zoom) Default() {
	*z = Zoom(2)
}
