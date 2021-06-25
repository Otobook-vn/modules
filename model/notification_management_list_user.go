package model

// NotificationManagementListUser ...
type NotificationManagementListUser struct {
	ID string `gorm:"primaryKey"`

	// Ref
	NotificationManagement   NotificationManagement `gorm:"foreignKey:NotificationManagementID"`
	NotificationManagementID string
	User                     User `gorm:"foreignKey:UserID"`
	UserID                   string
}
