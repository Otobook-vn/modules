package model

import (
	"time"
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

// IsStatus ...
func (n NotificationManagement) IsStatus(status string) bool {
	return n.Status == status
}

// IsTarget ...
func (n NotificationManagement) IsTarget(target string) bool {
	return n.Target == target
}
