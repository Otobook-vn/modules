package model

import "time"

// NotificationManagementListUser ...
type NotificationManagementListUser struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time

	// Ref
	NotificationManagement   NotificationManagement `gorm:"foreignKey:NotificationManagementID"`
	NotificationManagementID string
	User                     User `gorm:"foreignKey:UserID"`
	UserID                   string
}

// TableName overrides the table name
func (NotificationManagementListUser) TableName() string {
	return "notification_managements_list_users"
}
