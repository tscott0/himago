package himago

import "math"

// SatTime is a defines time to minute precision.
type SatTime struct {
	Year   *int
	Month  *int
	Day    *int
	Hour   *int
	Minute *int
}

// Round down do the nearest 10 minute multiple and update
// the Minute value.  e.g. at 13:34 would become 13:30
// Existing multiples of 10 minutes will be affected.
func (time *SatTime) Round() {
	newMins := int(10 * math.Floor(float64(*time.Minute)/10))

	time.Minute = &newMins
}

// TODO: func to roll back 10 mins
