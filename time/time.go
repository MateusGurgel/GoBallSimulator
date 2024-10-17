package time

import "time"

type Time struct {
	DeltaTime  float64
	lastUpdate time.Time
}

func NewTime() *Time {
	newTime := Time{}
	newTime.UpdateDeltaTime()
	return &newTime
}

func calculateDeltaTime(lastUpdate time.Time) float64 {

	if lastUpdate.IsZero() {
		return 0.0
	}

	now := time.Now()
	return now.Sub(lastUpdate).Seconds()
}

func (t *Time) UpdateDeltaTime() {
	now := time.Now()
	t.DeltaTime = calculateDeltaTime(t.lastUpdate)
	t.lastUpdate = now
}
