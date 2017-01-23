package himago

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Xy represents coordinates or dimensions of an image.
// X and Y should always be positive integers.
type Xy struct {
	X int
	Y int
}

// Outputs Xy as two integers separated by an "x" e.g. 100x40
// This is the same format that Set accepts as input.
func (c *Xy) String() string {
	return fmt.Sprintf("%vx%v", c.X, c.Y)
}

// Set accepts two integers separated by an "x" e.g. 100x40
// It then sets the values of X and Y.
// Implements the flag.Value interface.
func (c *Xy) Set(value string) error {
	value = strings.ToLower(value)
	coords := strings.Split(value, "x")

	if len(coords) != 2 {
		return errors.New("Invalid coordinates.")
	}

	var err error
	c.X, err = strconv.Atoi(coords[0])
	if err != nil {
		return errors.New("X was an invalid number. Received: " + value)
	}

	c.Y, err = strconv.Atoi(coords[1])
	if err != nil {
		return errors.New("Y was an invalid number. Received: " + value)
	}

	if c.X < 0 || c.Y < 0 {
		return errors.New("Coordinates cannot be negative.")
	}

	return nil
}
