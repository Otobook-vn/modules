package model

import (
	"time"

	"gorm.io/datatypes"
)

// UserNotification ...
type UserNotification struct {
	ID        string
	Category  string
	Type      string
	TargetID  string
	Data      datatypes.JSON // Custom format for each type of notification
	IsRead    bool
	CreatedAt time.Time

	// Ref
	User   User `gorm:"foreignKey:UserID"`
	UserID string
}

// TableName overrides the table name
func (UserNotification) TableName() string {
	return "user_notifications"
}

// UserNotificationData ...
type UserNotificationData struct {
	Status   string `json:"status,omitempty"`
	UserID   string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

// UserNotificationAction ...
type UserNotificationAction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// UserNotificationPushData ...
type UserNotificationPushData struct {
	ID      string                  `json:"id"`
	Title   string                  `json:"title"`
	Content string                  `json:"content"`
	Action  *UserNotificationAction `json:"action"`
	Photo   string                  `json:"photo"`
}

// UserNotificationResponse ...
type UserNotificationResponse struct {
	ID      string                  `json:"id"`
	Title   string                  `json:"title"`
	Content string                  `json:"content"`
	Action  *UserNotificationAction `json:"action"`
	Photo   string                  `json:"photo"`
	Data    *UserNotificationData   `json:"-"`
}
