package himago

import (
	"fmt"
	"math"
	"time"
)

// SatTime is a defines time to minute precision.
// Just wraps time.Time
type SatTime struct {
	time.Time
}

// Round down do the nearest 10 minute multiple and update
// the Minute value.  e.g. at 13:34 would become 13:30
// Existing multiples of 10 minutes will not be affected.
func (t *SatTime) Round() {
	newMins := int(10 * math.Floor(float64(t.Minute())/10))
	fmt.Printf("Rounding to %v\n", newMins)

	t.Time = time.Date(t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		newMins,
		0, 0, time.UTC)

	fmt.Printf("Rounded to %v\n", t.Minute())

}

// Go back 10 minutes. Assumes that the time was previously a
// multiple of 10 minutes, which calling Round() guarantees.
func (t *SatTime) Rollback() {
	t.Time = t.Add(-10 * time.Minute)
}
