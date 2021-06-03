package utils

import "time"

const timezoneHCM = "Asia/Ho_Chi_Minh"

// getHCMLocation ...
func getHCMLocation() *time.Location {
	l, _ := time.LoadLocation(timezoneHCM)
	return l
}

// StartOfDay ...
func StartOfDay(t time.Time) time.Time {
	loc := getHCMLocation()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}

// TimeBeforeNowInMin ...
func TimeBeforeNowInMin(min int) time.Time {
	loc := getHCMLocation()
	return time.Now().Add(time.Minute * time.Duration(min) * -1).In(loc)
}

// TimeNow ...
func TimeNow() time.Time {
	loc := getHCMLocation()
	return time.Now().In(loc)
}
