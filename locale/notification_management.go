package locale

// Constant
const (
	NotificationManagementKeyInvalidTitle       = "notification_management_invalid_title"
	NotificationManagementKeyInvalidCategory    = "notification_management_invalid_category"
	NotificationManagementKeyInvalidUserGroup   = "notification_management_invalid_user_group"
	NotificationManagementKeyInvalidExternalURL = "notification_management_invalid_external_url"
	NotificationManagementKeyInvalidContent     = "notification_management_invalid_content"
	NotificationManagementKeyInvalidSendOption  = "notification_management_invalid_send_option"
	NotificationManagementKeyListUserEmpty      = "notification_management_list_user_empty"
	NotificationManagementKeyCannotSend         = "notification_management_cannot_send"
	NotificationManagementKeyInvalidSender      = "notification_management_invalid_sender"
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
		Key:     NotificationManagementKeyInvalidExternalURL,
		Message: "liên kết ngoài không hợp lệ",
		Code:    503,
	},
	{
		Key:     NotificationManagementKeyInvalidContent,
		Message: "nội dung tin nhắn không hợp lệ",
		Code:    504,
	},
	{
		Key:     NotificationManagementKeyInvalidSendOption,
		Message: "loại người nhận không hợp lệ",
		Code:    505,
	},
	{
		Key:     NotificationManagementKeyListUserEmpty,
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
}
