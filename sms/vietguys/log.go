package vietguys

import (
	"fmt"
	"time"

	"github.com/Otobook-vn/modules/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	otpValidMinute = 30
)

// Log ...
type Log struct {
	ID          primitive.ObjectID `bson:"_id"`
	Carrier     string             `bson:"carrier"`
	Type        string             `bson:"type"`
	PhoneNumber string             `bson:"phoneNumber"`
	Code        string             `bson:"code"`
	IsCodeValid bool               `bson:"isCodeValid"`
	Content     string             `bson:"content"`
	IP          string             `bson:"ip"`
	Success     bool               `bson:"success"`
	Result      string             `bson:"result"`
	CreatedAt   time.Time          `bson:"createdAt"`

	tableName string
}

// Save log to db
func (s Service) saveLog(doc Log) {
	if _, err := s.DB.InsertOne(bgCtx, doc); err != nil {
		fmt.Println("*** Error when create log", err)
		fmt.Println("*** Log", doc)
	}
}

// Check phone or ip can send sms
func (s Service) checkCanSend(phone, ip string) bool {
	var (
		canSend      = true
		startOfToday = utils.StartOfDay(time.Now())
	)

	// Check phone number first
	if len(phone) > 0 && s.PhoneMaxSendPerDay > 0 {
		total, _ := s.DB.CountDocuments(bgCtx, bson.M{
			"phoneNumber": phone,
			"createdAt": bson.M{
				"$gte": startOfToday,
			},
		})
		canSend = total > int64(s.PhoneMaxSendPerDay)
	}

	// Check ip, but only check if can send
	if canSend && len(ip) > 0 && s.IPMaxSendPerDay > 0 {
		total, _ := s.DB.CountDocuments(bgCtx, bson.M{
			"ip": ip,
			"createdAt": bson.M{
				"$gte": startOfToday,
			},
		})
		canSend = total > int64(s.IPMaxSendPerDay)
	}

	return canSend
}

// Check otp right or not
func (s Service) checkOTP(phone, otpCode string) bool {
	var (
		timeAgo = utils.TimeBeforeNowInMin(otpValidMinute)
	)

	total, _ := s.DB.CountDocuments(bgCtx, bson.M{
		"phoneNumber": phone,
		"code":        otpCode,
		"isCodeValid": true,
		"createdAt": bson.M{
			"$gte": timeAgo,
		},
	})
	isValid := total > 0

	// If success, set code valid to false
	if isValid {
		s.DB.UpdateOne(bgCtx, bson.M{
			"code": otpCode,
		}, bson.M{
			"$set": bson.M{
				"isCodeValid": false,
			},
		})
	}
	return isValid
}
