package audit

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// newMongoID ...
func newMongoID() primitive.ObjectID {
	return primitive.NewObjectID()
}

const timezoneHCM = "Asia/Ho_Chi_Minh"

// getHCMLocation ...
func getHCMLocation() *time.Location {
	l, _ := time.LoadLocation(timezoneHCM)
	return l
}

// timeNow ...
func timeNow() time.Time {
	loc := getHCMLocation()
	return time.Now().In(loc)
}
