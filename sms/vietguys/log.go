package vietguys

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Log ...
type Log struct {
	ID        string `gorm:"primaryKey"`
	Carrier   string
	Type      string
	Recipient string
	Content   string
	Success   bool
	Result    string
	CreatedAt time.Time
}

// LogTable define table will save data
func LogTable(tableName string) func (tx *gorm.DB) *gorm.DB {
	return func (tx *gorm.DB) *gorm.DB {
		if tableName == "" {
			return tx.Table("vietguys_logs")
		}

		return tx.Table(tableName)
	}
}

// Save log to db
func (s Service) saveLog(doc Log) {
	if err := s.Config.PostgreSQL.Model(Log{}).Scopes(LogTable(s.Config.TableName)).Create(doc).Error; err != nil {
		fmt.Println("*** Error when create log", err)
		fmt.Println("*** Log", doc)
	}
}