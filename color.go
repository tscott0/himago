package himago

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

// Color wraps a color.NRGBA for passing as a command-line flag.
type Color struct {
	color.NRGBA
}

// String returns the color as a string
func (c *Color) String() string {
	return fmt.Sprintf("R%v G%v B%v A%v,", c.R, c.G, c.B, c.A)
}

// Set accepts a string of digits making a hex number.
// Must be prefixed with a #
// eg. #FFFFFF or #ab12ab
func (c *Color) Set(value string) error {

	if !strings.HasPrefix(value, "#") {
		return errors.New("Colours must start with a #")
	}

	chars := strings.Split(value, "")

	if len(chars) != 7 {
		return errors.New("Colour string is not the right length")
	}

	rString := strings.Join(chars[1:3], "")
	gString := strings.Join(chars[3:5], "")
	bString := strings.Join(chars[5:7], "")

	r, err := strconv.ParseUint(rString, 16, 32)
	if err != nil {
		return errors.New("Invalid hexadecimal number (red)")
	}

	g, err := strconv.ParseUint(gString, 16, 32)
	if err != nil {
		return errors.New("Invalid hexadecimal number (green)")
	}

	b, err := strconv.ParseUint(bString, 16, 32)
	if err != nil {
		return errors.New("Invalid hexadecimal number (blue)")
	}

	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)
	c.A = 255

	return nil
}
