package himago

import "testing"

// TestZoomString tests zoom.String() returns the zoom level as a string.
func TestZoomString(t *testing.T) {
	zoom := Zoom(0)
	zoomString := zoom.String()

	if zoomString != "0" {
		t.Errorf("Expected \"0\", received %v", zoomString)
	}
}

// TestZoomSet tests that a Zoom is correctly initialised
// from a command-line flag representing it.
// The flag.Value interface requires a Set() method.
func TestZoomSet(t *testing.T) {
	validZooms := []struct {
		name string
		in   string
		out  int
	}{
		{"Zoom 1", "1", 1},
		{"Zoom 01", "01", 1},
		{"Zoom 001", "001", 1},
		{"Zoom 5", "5", 5},
	}

	for _, vz := range validZooms {
		t.Run(vz.name, func(t *testing.T) {
			var zoom Zoom
			err := zoom.Set(vz.in)
			if err != nil {
				t.Errorf("Failed to call zoom.Set(\"%v\")", vz.in)
			}

			intZoom := int(zoom)
			expected := vz.out

			if intZoom != expected {
				t.Errorf("Expected \"%v\", received \"%v\"", expected, intZoom)
			}
		})
	}
}

// TestZoomSetInvalid tests that a Zoom does not allow calling
// Set() with invalid values.
func TestZoomSetInvalid(t *testing.T) {
	invalidZooms := []struct {
		name string
		in   string
	}{
		{"Zoom 0", "0"},
		{"Zoom 6", "6"},
		{"Zoom text", "text"},
		{"Zoom 5 with leading whitespace", " 5"},
		{"Zoom 5 with trailing whitespace", "5 "},
		{"Zoom whitespace", " "},
	}

	for _, iz := range invalidZooms {
		t.Run(iz.name, func(t *testing.T) {
			var zoom Zoom
			err := zoom.Set(iz.in)

			// Check that there was an error
			if err == nil {
				t.Errorf("Calling zoom.Set(\"%v\") should have thrown an error", iz.in)
			}
			// If Set() fails the Zoom will have the default value
		})
	}
}

// TestZoomGridWidth tests that a Zoom returns the expected width for a grid.
// This number represents the size of the grid in Tiles
func TestZoomGridWidth(t *testing.T) {
	validZooms := []struct {
		name string
		in   string
		out  int
	}{
		{"Zoom 1", "1", 1},
		{"Zoom 2", "2", 2},
		{"Zoom 3", "3", 4},
		{"Zoom 4", "4", 8},
		{"Zoom 5", "5", 16},
	}

	for _, vz := range validZooms {
		t.Run(vz.name, func(t *testing.T) {
			var zoom Zoom
			_ = zoom.Set(vz.in)

			width := zoom.GridWidth()

			if width != vz.out {
				t.Errorf("Expected \"%v\", received \"%v\"", vz.out, width)
			}
		})
	}
}

// TestZoomIsSet checks that the boolan IsSet() returns
// true when the Zoom is non-zero.
func TestZoomIsSetTrue(t *testing.T) {
	var zoom Zoom
	_ = zoom.Set("1")

	if !zoom.IsSet() {
		t.Errorf("Expected IsSet() to return true")
	}
}

// TestZoomIsSet checks that the boolan IsSet() returns
// false when the Zoom is zero.
func TestZoomIsSetFalse(t *testing.T) {
	var zoom Zoom

	if zoom.IsSet() {
		t.Errorf("Expected IsSet() to return false")
	}
}

// TestZoomDefault checks that the Zoom correctly initialises
// to 2 when calling Default()
func TestZoomDefault(t *testing.T) {
	expected := 2

	var zoom Zoom
	zoom.Default()

	if int(zoom) != expected {
		t.Errorf("Expected \"%v\", received \"%v\"", expected, zoom)
	}
}
