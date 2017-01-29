package himago

import (
	"errors"
	"fmt"
	"strings"
)

// BandURL is a Sprintf formatted string representing the URL
// needed to Get images in the specified light band.
// Valid values include "standard" and numbers 01-16 and are defined
// in the baseURLMap below.
type BandURL string

var (
	uRLPrefix string = "http://himawari8-dl.nict.go.jp/himawari8/img/"
	uRLSuffix string = "/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"

	baseURLMap = map[string]string{
		"standard": uRLPrefix + "D531106" + uRLSuffix,
		"01":       uRLPrefix + "FULL_24h/B01" + uRLSuffix,
		"02":       uRLPrefix + "FULL_24h/B02" + uRLSuffix,
		"03":       uRLPrefix + "FULL_24h/B03" + uRLSuffix,
		"04":       uRLPrefix + "FULL_24h/B04" + uRLSuffix,
		"05":       uRLPrefix + "FULL_24h/B05" + uRLSuffix,
		"06":       uRLPrefix + "FULL_24h/B06" + uRLSuffix,
		"07":       uRLPrefix + "FULL_24h/B07" + uRLSuffix,
		"08":       uRLPrefix + "FULL_24h/B08" + uRLSuffix,
		"09":       uRLPrefix + "FULL_24h/B09" + uRLSuffix,
		"10":       uRLPrefix + "FULL_24h/B10" + uRLSuffix,
		"11":       uRLPrefix + "FULL_24h/B11" + uRLSuffix,
		"12":       uRLPrefix + "FULL_24h/B12" + uRLSuffix,
		"13":       uRLPrefix + "FULL_24h/B13" + uRLSuffix,
		"14":       uRLPrefix + "FULL_24h/B14" + uRLSuffix,
		"15":       uRLPrefix + "FULL_24h/B15" + uRLSuffix,
		"16":       uRLPrefix + "FULL_24h/B16" + uRLSuffix,
	}
)

// String returns BandURL as a string. Nothing to see here.
func (b *BandURL) String() string {
	return string(*b)
}

// Set will lookup a band and set BandURL for that band.
// e.g. a key of "03" will set the value to be
// "http://himawari8-dl.nict.go.jp/himawari8/img/FULL_24H/B03/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
// Assumes zero-padding for numbers. Passing "01" will work but "1" will not.
func (b *BandURL) Set(key string) error {
	// Ensure lookups are down with lowercase string
	// Only applies to "standard" for now
	key = strings.ToLower(key)

	fmt.Println("key: " + key)

	// Check that the value passed exists or return an error
	if val, ok := baseURLMap[key]; ok {
		bURL := BandURL(val)
		fmt.Println("val: " + val)
		*b = bURL
		fmt.Println("b:   " + string(*b))
	} else {
		return errors.New("Invalid band. Received \"" + key + "\"")
	}

	return nil
}

// IsSet returns true if the BandURL is the non-default value.
// If the underlying string is "" then return false.
func (b *BandURL) IsSet() bool {
	return string(*b) != ""
}
