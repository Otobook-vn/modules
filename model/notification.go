package model

import (
	"encoding/json"
	"strings"
	"time"

	"gorm.io/datatypes"

	"github.com/Otobook-vn/modules/constant"
	"github.com/Otobook-vn/modules/translation"
)

// Notification ...
type Notification struct {
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
func (Notification) TableName() string {
	return "notifications"
}

// NotificationData ...
type NotificationData struct {
	Status   string `json:"status,omitempty"`
	UserID   string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

// NotificationAction ...
type NotificationAction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// NotificationPushData ...
type NotificationPushData struct {
	ID      string              `json:"id"`
	Title   string              `json:"title"`
	Content string              `json:"content"`
	Action  *NotificationAction `json:"action"`
	Photo   string              `json:"photo"`
}

// NotificationResponse ...
type NotificationResponse struct {
	ID      string              `json:"id"`
	Title   string              `json:"title"`
	Content string              `json:"content"`
	Action  *NotificationAction `json:"action"`
	Photo   string              `json:"photo"`
	Data    *NotificationData   `json:"-"`
}

// ParseNotificationDataFromJSON ...
func ParseNotificationDataFromJSON(jsonData datatypes.JSON) (result *NotificationData) {
	if jsonData == nil {
		return
	}
	bytes, _ := json.Marshal(jsonData)
	if err := json.Unmarshal(bytes, &result); err != nil {
		return
	}
	return
}

// GetNotificationAction based on type
func GetNotificationAction(notificationType string, targetID string) (result NotificationAction) {
	switch notificationType {
	case constant.NotificationTypeConsultQuestionAdminChangeStatus,
		constant.NotificationTypeConsultQuestionSpecialistMarkCompleted,
		constant.NotificationTypeConsultQuestionHasNewComment,
		constant.NotificationTypeConsultQuestionUserRating:
		result.Type = constant.NotificationActionOpenConsultQuestionDetail
		result.Value = targetID
	}

	// Uppercase type
	result.Type = strings.ToUpper(result.Type)

	return
}

// GetNotificationTitle based on category and data
func GetNotificationTitle(lang string, notificationCategory string, data *NotificationData) string {
	return translation.GetNotificationTitle(lang, notificationCategory, data)
}

// GetNotificationContent based on type and data
func GetNotificationContent(lang string, notificationType string, data *NotificationData) string {
	// Get translated data first
	var translatedData = *data
	if data != nil {
		// NOTE: messageID must match with key declared in toml files /translation

		// Get by type
		switch notificationType {
		case constant.NotificationTypeConsultQuestionAdminChangeStatus:
			translatedData.Status = translation.GetOtherText(lang, "consult_question_status_"+data.Status)
		}
	}

	return translation.GetNotificationContent(lang, notificationType, translatedData)
}
