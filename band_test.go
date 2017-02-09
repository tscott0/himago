package himago

import "testing"

// TestString tests band.String() returns the band as a string.
// The flag.Value interface requires a String() method.
func TestBandString(t *testing.T) {
	band := Band(0)
	bandString := band.String()

	if bandString != "0" {
		t.Errorf("Expected \"0\", received %v", bandString)
	}
}

// TestBandURLDefault tests that the correct URL is returned
// when calling URL() if band has the default value (0)
func TestBandURLDefault(t *testing.T) {
	band := Band(0)

	if band.URL() != "http://himawari8-dl.nict.go.jp/himawari8/img/D531106/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png" {
		t.Errorf("Did not return default URL")
	}
}

// TestBandURL tests that the correct URL is returned
// when calling URL() if band has non-default values.
func TestBandURL(t *testing.T) {
	// We are constructing a URL with zero-padding.
	// Testing 1 to verify it works for single-digit bands.
	t.Run("Band 1", func(t *testing.T) {
		band := Band(1)
		expected := "http://himawari8-dl.nict.go.jp/himawari8/img/FULL_24h/B01/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
		url := band.URL()

		if url != expected {
			t.Errorf("Failed to construct URL for band \"%v\"\nExpected: \"%v\"\nReceived: \"%v\"", band, expected, url)
		}
	})

	// Testing 16 (max) to verify it works for two-digit bands
	t.Run("Band 16", func(t *testing.T) {
		band := Band(16)
		expected := "http://himawari8-dl.nict.go.jp/himawari8/img/FULL_24h/B16/%vd/550/%02d/%02d/%02d/%02d%02d00_%v_%v.png"
		url := band.URL()

		if url != expected {
			t.Errorf("Failed to construct URL for band \"%v\"\nExpected: \"%v\"\nReceived: \"%v\"", band, expected, url)
		}
	})
}

// TestBandSet tests that a Band is correctly initialised
// from a command-line flag representing it.
// The flag.Value interface requires a Set() method.
func TestBandSet(t *testing.T) {
	validBands := []struct {
		name string
		in   string
		out  int
	}{
		{"Band 1", "1", 1},
		{"Band 01", "01", 1},
		{"Band 001", "001", 1},
		{"Band 16", "16", 16},
	}

	for _, vb := range validBands {
		t.Run(vb.name, func(t *testing.T) {
			var band Band
			err := band.Set(vb.in)
			if err != nil {
				t.Errorf("Failed to call band.Set(\"%v\")", vb.in)
			}

			intBand := int(band)
			expected := vb.out

			if intBand != expected {
				t.Errorf("Expected \"%v\", received \"%v\"", expected, intBand)
			}
		})
	}
}

// TestBandSetInvalid tests that a Band does not allow calling
// Set() with invalid values.
func TestBandSetInvalid(t *testing.T) {
	invalidBands := []struct {
		name string
		in   string
	}{
		{"Band 0", "0"},
		{"Band 17", "17"},
		{"Band text", "text"},
		{"Band 5 with leading whitespace", " 5"},
		{"Band 5 with trailing whitespace", "5 "},
		{"Band whitespace", " "},
	}

	for _, ib := range invalidBands {
		t.Run(ib.name, func(t *testing.T) {
			var band Band
			err := band.Set(ib.in)

			// Check that there was an error
			if err == nil {
				t.Errorf("Calling band.Set(\"%v\") should have thrown an error", ib.in)
			}
			// If Set() fails the Band will have the default value
		})
	}
}
