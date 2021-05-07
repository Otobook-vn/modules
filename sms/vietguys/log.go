package vietguys

import (
	"fmt"
	"time"

	"github.com/Otobook-vn/modules/utils"
	"gorm.io/gorm"
)

// Log ...
type Log struct {
	ID          string `gorm:"primaryKey"`
	Carrier     string
	Type        string
	PhoneNumber string
	Content     string
	IP          string
	Success     bool
	Result      string
	CreatedAt   time.Time
}

// LogTable define table will save data
func LogTable(tableName string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if tableName == "" {
			return tx.Table("vietguys_logs")
		}

		return tx.Table(tableName)
	}
}

// Save log to db
func (s Service) saveLog(doc Log) {
	if err := s.PostgreSQL.Model(Log{}).Scopes(LogTable(s.LogTableName)).Create(doc).Error; err != nil {
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
	if s.PhoneMaxSendPerDay > 0 {
		s.PostgreSQL.Model(Log{}).Where("phone_number = ? AND created_at >=", phone, startOfToday).Count(&count)
		canSend = count > int64(s.PhoneMaxSendPerDay)
	}

	// Check ip, but only check if can send
	if canSend && s.IPMaxSendPerDay > 0 {
		s.PostgreSQL.Model(Log{}).Where("ip = ? AND created_at >=", ip, startOfToday).Count(&count)
		canSend = count > int64(s.IPMaxSendPerDay)
	}

	return canSend
}
