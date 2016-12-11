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
func (coord *Xy) String() string {
	return fmt.Sprintf("%vx%v", coord.X, coord.Y)
}

// Set accepts two integers separated by an "x" e.g. 100x40
// It then sets the values of X and Y.
// Implements the flag.Value interface.
func (coord *Xy) Set(value string) error {
	value = strings.ToLower(value)
	coords := strings.Split(value, "x")

	if len(coords) != 2 {
		return errors.New("Invalid coordinates.")
	}

	var err error
	coord.X, err = strconv.Atoi(coords[0])
	if err != nil {
		return errors.New("X was an invalid number. Received: " + value)
	}

	coord.Y, err = strconv.Atoi(coords[1])
	if err != nil {
		return errors.New("Y was an invalid number. Received: " + value)
	}

	if coord.X < 0 || coord.Y < 0 {
		return errors.New("Coordinates cannot be negative.")
	}

	return nil
}
