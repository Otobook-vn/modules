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
	IsCodeValid bool
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
	return "log_sms_vietguys"
}

// Save log to db
func (s Service) saveLog(doc Log) {
	// Return if no postgresql instance
	if s.PostgreSQL == nil {
		return
	}

	if err := s.PostgreSQL.Model(Log{}).Create(doc).Error; err != nil {
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
		timeAgo       = utils.TimeBeforeNowInMin(otpValidMinute)
		count   int64 = 0
	)

	s.PostgreSQL.Debug().Model(Log{}).Where("phone_number = ? AND code = ? AND is_code_valid = ? AND created_at >= ?", phone, otpCode, true, timeAgo).Count(&count)
	isValid := count > 0
	// If success, set code valid to false
	if isValid {
		s.PostgreSQL.Model(Log{}).Where("code = ?", otpCode).Update("is_code_valid", false)
	}
	return isValid
}
