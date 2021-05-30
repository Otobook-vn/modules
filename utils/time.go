package utils

import "time"

// StartOfDay ...
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UTC()
}

// TimeBeforeNowInMin ...
func TimeBeforeNowInMin(min int) time.Time {
	return time.Now().Add(time.Minute * time.Duration(min) * -1).UTC()
}

// TimeNowUTC ...
func TimeNowUTC() time.Time {
	return time.Now().UTC()
}
