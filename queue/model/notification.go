package queuemodel

import (
	"github.com/Otobook-vn/modules/model"
)

// NotificationCreateAndPushSingle ...
type NotificationCreateAndPushSingle struct {
	Notification model.Notification
	IsPush       bool
}

// NotificationAdminPushAll ...
type NotificationAdminPushAll struct {
	Data model.NotificationManagement
}

// NotificationAdminPushByListUsers ...
type NotificationAdminPushByListUsers struct {
	Data model.NotificationManagement
}
