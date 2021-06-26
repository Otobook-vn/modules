package model

import (
	"time"

	"gorm.io/datatypes"
)

// NotificationManagement ...
type NotificationManagement struct {
	ID          string `gorm:"primaryKey"`
	Language    string // vi, en
	Category    string // common, consult
	UserGroup   string // all, normal, specialist
	Title       string
	Content     string
	ActionType  string
	ActionValue string
	ActionData  datatypes.JSON
	Target      string // all_users, by_list
	Status      string // pending, sent
	StatsSent   int64  // total sent
	CreatedAt   time.Time
	SentAt      time.Time

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
