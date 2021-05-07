package utils

import "time"

// StartOfDay ...
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UTC()
}

// MinBeforeNow ...
func MinBeforeNow(min int) time.Time {
	return time.Now().Add(time.Minute * time.Duration(min)).UTC()
}
