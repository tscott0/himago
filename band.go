package himago

import (
	"errors"
	"fmt"
	"strconv"
)

// Band is an int representing the electromagnetic frequency that an image was
// taken.
//
//   Band  Frequency  Description
//   1     00.47µm    BLUE
//   2     00.51µm    GREEN
//   3     00.64µm    RED
//   4     00.86µm    Near-IR
//   5     01.60µm    Near-IR
//   6     02.30µm    Near-IR
//   7     03.90µm    Short-IR
//   8     06.20µm    Mid-IR
//   9     06.90µm    Mid-IR
//   10    07.30µm    Mid-IR
//   11    08.60µm    Far-IR
//   12    09.60µm    Far-IR
//   13    10.40µm    Far-IR
//   14    11.20µm    Far-IR
//   15    12.40µm    Far-IR
//   16    13.30µm    Far-IR
//
// 0 represents the default band which is a full-colour version
// combining the visible light bands (RGB)
type Band int

var (
	uRLPrefix  = "http://himawari8-dl.nict.go.jp/himawari8/img/"
	uRLSuffix  = "/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
	defaultURL = uRLPrefix + "D531106" + uRLSuffix
)

// String returns the Band int as a string. Nothing to see here.
func (b *Band) String() string {
	return fmt.Sprintf("%v", *b)
}

// URL will build a URL string for the band
// The URL will be in Sprintf format e.g. a band
// of "3" will set the value to be:
//     "http://himawari8-dl.nict.go.jp/himawari8/img/FULL_24H/B03/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
func (b *Band) URL() string {

	if int(*b) == 0 {
		return defaultURL
	} else {
		// Construct the full URL and return it
		return fmt.Sprintf("%sFULL_24h/B%02d%s", uRLPrefix, *b, uRLSuffix)
	}
}

// Set will take the flag passed as a string and attempt to convert it
// to an int. That int is used the set the value of the Band.
func (b *Band) Set(flag string) error {
	// Attempt to cast to int
	band, err := strconv.Atoi(flag)

	// If it's not an integer or isn't between 1 and 16 (inclusive) error
	if err != nil || band < 1 || band > 16 {
		return errors.New("Band must be an integer between 1 and 16 inclusive")
	}

	// Set the value of the band
	*b = Band(band)

	return nil
}
