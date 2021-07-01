package locale

// Constant
const (
	UserKeyInvalidName             = "user_invalid_name"
	UserKeyInvalidPhoneCountryCode = "user_invalid_phone_country_code"
	UserKeyInvalidPhoneNumber      = "user_invalid_phone_number"
	UserKeyInvalidEmail            = "user_invalid_email"
	UserKeyInvalidPassword         = "user_invalid_password"
	UserKeyInvalidReferralCode     = "user_invalid_referral_code"
	UserKeyPhoneExisted            = "user_phone_existed"
	UserKeyNotFound                = "user_not_found"
	UserKeyInvalidOldPassword      = "user_invalid_old_password"
	UserKeyInvalidNewPassword      = "user_invalid_new_password"
	UserKeyWrongOldPassword        = "user_wrong_old_password"
	UserKeyInvalidDeviceID         = "user_invalid_device_id"
	UserKeyInvalidOSName           = "user_invalid_os_name"
	UserKeyInvalidOSVersion        = "user_invalid_os_version"
	UserKeyInvalidAppVersion       = "user_invalid_app_version"
	UserKeyInvalidGroup            = "user_invalid_group"
	UserKeyInvalidRegisterFrom     = "user_invalid_register_from"
	UserKeyCodeExisted             = "user_code_existed"
	UserKeyEmailExisted            = "user_email_existed"
	UserKeyReferralCodeExisted     = "user_referral_code_existed"
)

// 250-299
var user = []Locale{
	{
		Key:     UserKeyInvalidName,
		Message: "tên phải từ 3 ký tự",
		Code:    250,
	},
	{
		Key:     UserKeyInvalidPhoneCountryCode,
		Message: "mã quốc gia không hợp lệ",
		Code:    251,
	},
	{
		Key:     UserKeyInvalidPhoneNumber,
		Message: "số điện thoại không hợp lệ",
		Code:    252,
	},
	{
		Key:     UserKeyInvalidEmail,
		Message: "email không hợp lệ",
		Code:    253,
	},
	{
		Key:     UserKeyInvalidPassword,
		Message: "mật khẩu phải từ 6 ký tự",
		Code:    254,
	},
	{
		Key:     UserKeyPhoneExisted,
		Message: "số điện thoại đã tồn tại trong hệ thống",
		Code:    255,
	},
	{
		Key:     UserKeyInvalidReferralCode,
		Message: "mã mời bạn không hợp lệ",
		Code:    256,
	},
	{
		Key:     UserKeyNotFound,
		Message: "người dùng không tìm thấy",
		Code:    257,
	},
	{
		Key:     UserKeyInvalidOldPassword,
		Message: "mật khẩu cũ phải từ 6 ký tự",
		Code:    258,
	},
	{
		Key:     UserKeyInvalidNewPassword,
		Message: "mật khẩu mới phải từ 6 ký tự",
		Code:    259,
	},
	{
		Key:     UserKeyWrongOldPassword,
		Message: "mật khẩu cũ không chính xác",
		Code:    260,
	},
	{
		Key:     UserKeyInvalidDeviceID,
		Message: "mã thiết bị không hợp lệ",
		Code:    261,
	},
	{
		Key:     UserKeyInvalidOSName,
		Message: "tên OS không hợp lệ",
		Code:    262,
	},
	{
		Key:     UserKeyInvalidOSVersion,
		Message: "phiên bản OS không hợp lệ",
		Code:    263,
	},
	{
		Key:     UserKeyInvalidAppVersion,
		Message: "phiên bản app không hợp lệ",
		Code:    264,
	},
	{
		Key:     UserKeyInvalidGroup,
		Message: "nhóm người dùng không hợp lệ",
		Code:    265,
	},
	{
		Key:     UserKeyInvalidRegisterFrom,
		Message: "nguồn đăng ký không hợp lệ",
		Code:    266,
	},
	{
		Key:     UserKeyCodeExisted,
		Message: "mã khách hàng đã tồn tại",
		Code:    267,
	},
	{
		Key:     UserKeyEmailExisted,
		Message: "email đã tồn tại",
		Code:    268,
	},
	{
		Key:     UserKeyReferralCodeExisted,
		Message: "mã giới thiệu đã tồn tại",
		Code:    269,
	},
}
