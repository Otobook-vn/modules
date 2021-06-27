package locale

// Constant
const (
	NotificationManagementKeyInvalidTitle       = "notification_management_invalid_title"
	NotificationManagementKeyInvalidCategory    = "notification_management_invalid_category"
	NotificationManagementKeyInvalidUserGroup   = "notification_management_invalid_user_group"
	NotificationManagementKeyInvalidAction      = "notification_management_invalid_action"
	NotificationManagementKeyInvalidContent     = "notification_management_invalid_content"
	NotificationManagementKeyInvalidTarget      = "notification_management_invalid_target"
	NotificationManagementKeyListUsersEmpty     = "notification_management_list_user_empty"
	NotificationManagementKeyCannotSend         = "notification_management_cannot_send"
	NotificationManagementKeyInvalidSender      = "notification_management_invalid_sender"
	NotificationManagementKeyInvalidActionValue = "notification_management_invalid_action_value"
)

// 500-549
var notificationManagement = []Locale{
	{
		Key:     NotificationManagementKeyInvalidTitle,
		Message: "tiêu đề tin nhắn không hợp lệ",
		Code:    500,
	},
	{
		Key:     NotificationManagementKeyInvalidCategory,
		Message: "danh mục tin nhắn không hợp lệ",
		Code:    501,
	},
	{
		Key:     NotificationManagementKeyInvalidUserGroup,
		Message: "nhóm người dùng không hợp lệ",
		Code:    502,
	},
	{
		Key:     NotificationManagementKeyInvalidAction,
		Message: "hành động không hợp lệ",
		Code:    503,
	},
	{
		Key:     NotificationManagementKeyInvalidContent,
		Message: "nội dung tin nhắn không hợp lệ",
		Code:    504,
	},
	{
		Key:     NotificationManagementKeyInvalidTarget,
		Message: "loại người nhận không hợp lệ",
		Code:    505,
	},
	{
		Key:     NotificationManagementKeyListUsersEmpty,
		Message: "danh sách khách hàng không được trống",
		Code:    506,
	},
	{
		Key:     NotificationManagementKeyCannotSend,
		Message: "tin nhắn hiện không thể gửi",
		Code:    507,
	},
	{
		Key:     NotificationManagementKeyInvalidSender,
		Message: "người gửi không hợp lệ",
		Code:    508,
	},
	{
		Key:     NotificationManagementKeyInvalidActionValue,
		Message: "giá trị của hành động không hợp lệ",
		Code:    509,
	},
}
