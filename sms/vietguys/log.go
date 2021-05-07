package vietguys

import (
	"fmt"
	"time"

	"github.com/Otobook-vn/modules/utils"
)

const (
	otpValidMinute = 30
)

// Log ...
type Log struct {
	ID          string `gorm:"primaryKey"`
	Carrier     string
	Type        string
	PhoneNumber string
	Code        string
	Content     string
	IP          string
	Success     bool
	Result      string
	CreatedAt   time.Time

	tableName string
}

// TableName ...
func (l Log) TableName() string {
	if l.tableName != "" {
		return l.tableName
	}
	return "vietguys_logs"
}

// Save log to db
func (s Service) saveLog(doc Log) {
	if err := s.PostgreSQL.Model(Log{}).Scopes().Create(doc).Error; err != nil {
		fmt.Println("*** Error when create log", err)
		fmt.Println("*** Log", doc)
	}
}

// Check phone or ip can send sms
func (s Service) checkCanSend(phone, ip string) bool {
	var (
		canSend            = true
		count        int64 = 0
		startOfToday       = utils.StartOfDay(time.Now())
	)

	// Check phone number first
	if len(phone) > 0 && s.PhoneMaxSendPerDay > 0 {
		s.PostgreSQL.Model(Log{}).Where("phone_number = ? AND created_at >= ?", phone, startOfToday).Count(&count)
		canSend = count > int64(s.PhoneMaxSendPerDay)
	}

	// Check ip, but only check if can send
	if canSend && len(ip) > 0 && s.IPMaxSendPerDay > 0 {
		s.PostgreSQL.Model(Log{}).Where("ip = ? AND created_at >= ?", ip, startOfToday).Count(&count)
		canSend = count > int64(s.IPMaxSendPerDay)
	}

	return canSend
}

// Check otp right or not
func (s Service) checkOTP(phone, otpCode string) bool {
	var (
		minAgo       = utils.MinBeforeNow(otpValidMinute)
		count  int64 = 0
	)

	s.PostgreSQL.Model(Log{}).Where("phone_number = ? AND code = ? AND created_at >= ?", phone, otpCode, minAgo).Count(&count)
	return count > 0
}
