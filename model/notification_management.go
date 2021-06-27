package model

import (
	"time"

	"github.com/Otobook-vn/modules/constant"
)

// NotificationManagement ...
type NotificationManagement struct {
	ID                  string `gorm:"primaryKey"`
	Language            string // vi, en
	Category            string // common, consult
	Title               string
	Content             string
	ActionType          string
	ActionValue         string
	Target              string // all_users, by_list
	Status              string // pending, sent
	CreatedAt           time.Time
	SentAt              time.Time
	StatsSent           int
	StatsOpenInApp      int
	StatsOpenFromSystem int

	// Ref
	Creator   Staff `gorm:"foreignKey:CreatorID"`
	CreatorID string
	Sender    Staff `gorm:"foreignKey:SenderID"`
	SenderID  *string
}

// TableName overrides the table name
func (NotificationManagement) TableName() string {
	return "notification_managements"
}

// IsSent ...
func (n NotificationManagement) IsSent() bool {
	return n.Status == constant.NotificationSendStatusSent
}
