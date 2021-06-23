package queuemodel

import (
	"firebase.google.com/go/v4/messaging"
	"github.com/Otobook-vn/modules/model"
)

// NotificationCreateAndPushSingle ...
type NotificationCreateAndPushSingle struct {
	Notification model.Notification
	IsPush       bool
}

// NotificationPushByTokens ...
type NotificationPushByTokens struct {
	BatchID string
	Tokens  []string
	Message messaging.Message
}
