package himago

import (
	"testing"
	"time"
)

var roundTests = []struct {
	name string
	in   time.Time
	out  time.Time
}{
	{
		"Round down",
		time.Date(2016, time.Month(06), 20, 12, 45, 0, 0, time.UTC),
		time.Date(2016, time.Month(06), 20, 12, 40, 0, 0, time.UTC),
	},
	{
		"Round down edgecase",
		time.Date(2016, time.Month(06), 20, 12, 49, 59, 999999, time.UTC),
		time.Date(2016, time.Month(06), 20, 12, 40, 0, 0, time.UTC),
	},
	{
		"No change",
		time.Date(2016, time.Month(06), 20, 12, 00, 0, 0, time.UTC),
		time.Date(2016, time.Month(06), 20, 12, 00, 0, 0, time.UTC),
	},
}

// Loop over roundTests and wrap in inside a SatTime struct.
// Call in.Round() and fail if the underlying time.Time does not match out
func TestRound(t *testing.T) {
	for _, rt := range roundTests {
		// Start a subtest for each struct
		t.Run(rt.name, func(t *testing.T) {
			in := SatTime{rt.in}
			out := SatTime{rt.out}

			in.Round()

			if !in.Equal(out.Time) {
				t.Errorf("Received \"%v\", expected \"%v\"", in, out)
			}
		})
	}
}

var rollbackTests = []struct {
	name string
	in   time.Time
	out  time.Time
}{
	{
		"Rollback within hour",
		time.Date(2016, time.Month(06), 20, 12, 40, 0, 0, time.UTC),
		time.Date(2016, time.Month(06), 20, 12, 30, 0, 0, time.UTC),
	},
	{
		"Rollback hour",
		time.Date(2016, time.Month(06), 20, 12, 00, 0, 0, time.UTC),
		time.Date(2016, time.Month(06), 20, 11, 50, 0, 0, time.UTC),
	},
	{
		"Rollback year",
		time.Date(2016, time.Month(01), 01, 00, 00, 0, 0, time.UTC),
		time.Date(2015, time.Month(12), 31, 23, 50, 0, 0, time.UTC),
	},
}

// Loop over rollbackTests and wrap in inside a SatTime struct.
// Call in.Rollback() and fail if the underlying time.Time does not match out
// Test assumes Round() has already been called, meaning times are already in
// multiples of 10 minutes.
func TestRollback(t *testing.T) {
	for _, rt := range rollbackTests {
		// Start a subtest for each struct
		t.Run(rt.name, func(t *testing.T) {
			in := SatTime{rt.in}
			out := SatTime{rt.out}

			in.Rollback()

			if !in.Equal(out.Time) {
				t.Errorf("Received \"%v\", expected \"%v\"", in, out)
			}
		})
	}
}
