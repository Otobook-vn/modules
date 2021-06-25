package model

// NotificationManagement ...
type NotificationManagement struct {
}

// TableName overrides the table name
func (NotificationManagement) TableName() string {
	return "notification_managements"
}
