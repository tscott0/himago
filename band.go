package himago

import (
	"errors"
	"fmt"
	"strconv"
)

// BandURL is a Sprintf formatted string representing the URL
// needed to Get images in the specified light band.
// Valid values include numbers 1-16
type BandURL string

var (
	uRLPrefix  = "http://himawari8-dl.nict.go.jp/himawari8/img/"
	uRLSuffix  = "/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
	defaultURL = uRLPrefix + "D531106" + uRLSuffix
)

// String returns BandURL as a string. Nothing to see here.
func (b *BandURL) String() string {
	return string(*b)
}

// Set will lookup a band and set BandURL for that band.
// e.g. a band of "3" will set the value to be
// "http://himawari8-dl.nict.go.jp/himawari8/img/FULL_24H/B03/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
func (b *BandURL) Set(flag string) error {
	errMsg := "Band must be an integer between 1 and 16 inclusive"

	// Attempt to cast to int
	band, err := strconv.Atoi(flag)

	// If it's not an integer or isn't between 1 and 5 (inclusive) error
	if err != nil || band < 1 || band > 16 {
		return errors.New(errMsg)
	}

	// Create string with zero-padding e.g. FULL_24h/B12
	bandString := fmt.Sprintf("FULL_24h/B%02d", band)

	// Construct the full string
	*b = BandURL(uRLPrefix + bandString + uRLSuffix)

	return nil
}

// IsSet returns true if the BandURL is the non-default value.
// If the underlying string is "" then return false.
func (b *BandURL) IsSet() bool {
	return string(*b) != ""
}

// Default sets the BandURL to the default value
// Calling Default is required IsSet returns false
func (b *BandURL) Default() {
	*b = BandURL(defaultURL)
}
