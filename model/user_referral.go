package model

import "time"

// UserReferral ...
type UserReferral struct {
	ID        string `gorm:"primaryKey"`
	Device    string
	CreatedAt time.Time

	// Ref
	User      User `gorm:"foreignKey:UserID"`
	UserID    string
	Invitee   User `gorm:"foreignKey:InviteeID"`
	InviteeID string
}

// TableName overrides the table name
func (UserReferral) TableName() string {
	return "user_referrals"
}
