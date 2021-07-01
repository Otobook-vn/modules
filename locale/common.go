package locale

// Constant
const (
	CommonKeySuccess         = "common_success"
	CommonKeyBadRequest      = "common_bad_request"
	CommonKeyUnauthorized    = "common_unauthorized"
	CommonKeyNotFound        = "common_not_found"
	CommonKeyInvalidChecksum = "common_invalid_checksum"

	CommonKeyLoginWithUsernameFailed               = "common_login_with_username_failed"
	CommonKeyInvalidDeviceID                       = "common_invalid_device_id"
	CommonKeyEachDeviceCanInputReferralCodeOneTime = "common_each_device_can_input_referral_code_one_time"
	CommonKeyInvalidOTPCode                        = "common_invalid_otp_code"
	CommonKeyExpiredOTPCode                        = "common_expired_otp_code"
	CommonKeyInvalidCountry                        = "common_invalid_country"
	CommonKeyInvalidStatus                         = "common_invalid_status"
	CommonKeyInvalidCarBrandID                     = "common_invalid_car_brand_id"
	CommonKeyInvalidCarModelID                     = "common_invalid_car_model_id"
	CommonKeyInvalidCarModelVersionID              = "common_invalid_car_model_version_id"
	CommonKeyInvalidPagination                     = "common_invalid_pagination"
	CommonKeyInvalidProvince                       = "common_invalid_province"
	CommonKeyInvalidDistrict                       = "common_invalid_district"
	CommonKeyInvalidWard                           = "common_invalid_ward"
	CommonKeyInvalidSort                           = "common_invalid_sort"
	CommonKeyInvalidLanguage                       = "common_invalid_language"
	CommonKeyInvalidDateFrom                       = "common_invalid_date_from"
	CommonKeyInvalidDateTo                         = "common_invalid_date_to"
	CommonKeyInvalidPhoto                          = "common_invalid_photo"
	CommonKeyUserOtobookNotFound                   = "common_user_otobook_not_found"
	CommonKeyInvalidCarPartID                      = "common_invalid_car_part_id"
	CommonKeyInvalidUserID                         = "common_invalid_user_id"
	CommonKeyInvalidKeyword                        = "common_invalid_keyword"
	CommonKeyInvalidNotificationID                 = "common_invalid_notification_id"
	CommonKeErrorWhenUpdateData                    = "common_error_when_update_data"
)

// 1-99
var common = []Locale{
	{
		Key:     CommonKeySuccess,
		Message: "thành công",
		Code:    1,
	},
	{
		Key:     CommonKeyBadRequest,
		Message: "dữ liệu không hợp lệ",
		Code:    2,
	},
	{
		Key:     CommonKeyUnauthorized,
		Message: "bạn không có quyền thực hiện hành động này",
		Code:    3,
	},
	{
		Key:     CommonKeyNotFound,
		Message: "dữ liệu không tìm thấy",
		Code:    4,
	},
	{
		Key:     CommonKeyInvalidChecksum,
		Message: "xác thực dữ liệu thất bại",
		Code:    5,
	},
	{
		Key:     CommonKeyLoginWithUsernameFailed,
		Message: "tên đăng nhập hoặc mật khẩu không chính xác",
		Code:    6,
	},
	{
		Key:     CommonKeyInvalidDeviceID,
		Message: "mã tb không được trống",
		Code:    7,
	},
	{
		Key:     CommonKeyEachDeviceCanInputReferralCodeOneTime,
		Message: "mỗi thiết bị chỉ được nhập mã mời bạn 1 lần",
		Code:    8,
	},
	{
		Key:     CommonKeyInvalidOTPCode,
		Message: "mã otp không hợp lệ",
		Code:    9,
	},
	{
		Key:     CommonKeyExpiredOTPCode,
		Message: "mã otp không chính xác hoặc đã hết hạn",
		Code:    10,
	},
	{
		Key:     CommonKeyInvalidCountry,
		Message: "quốc gia không hợp lệ",
		Code:    11,
	},
	{
		Key:     CommonKeyInvalidStatus,
		Message: "trạng thái không hợp lệ",
		Code:    12,
	},
	{
		Key:     CommonKeyInvalidCarBrandID,
		Message: "id hãng xe không hợp lệ",
		Code:    13,
	},
	{
		Key:     CommonKeyInvalidCarModelID,
		Message: "id dòng xe không hợp lệ",
		Code:    14,
	},
	{
		Key:     CommonKeyInvalidCarModelVersionID,
		Message: "id kiểu xe không hợp lệ",
		Code:    15,
	},
	{
		Key:     CommonKeyInvalidPagination,
		Message: "phân trang không hợp lệ",
		Code:    16,
	},
	{
		Key:     CommonKeyInvalidProvince,
		Message: "tỉnh thành không hợp lệ",
		Code:    17,
	},
	{
		Key:     CommonKeyInvalidDistrict,
		Message: "quận huyện không hợp lệ",
		Code:    18,
	},
	{
		Key:     CommonKeyInvalidWard,
		Message: "xã phường không hợp lệ",
		Code:    19,
	},
	{
		Key:     CommonKeyInvalidSort,
		Message: "sắp xếp không hợp lệ",
		Code:    20,
	},
	{
		Key:     CommonKeyInvalidLanguage,
		Message: "ngôn ngữ không hợp lệ",
		Code:    21,
	},
	{
		Key:     CommonKeyInvalidDateFrom,
		Message: "thời gian bắt đầu không hợp lệ",
		Code:    22,
	},
	{
		Key:     CommonKeyInvalidDateTo,
		Message: "thời gian kết thúc không hợp lệ",
		Code:    23,
	},
	{
		Key:     CommonKeyInvalidPhoto,
		Message: "hình ảnh không hợp lệ",
		Code:    24,
	},
	{
		Key:     CommonKeyUserOtobookNotFound,
		Message: "người dùng Otobook không tìm thấy trong hệ thống, vui lòng khởi tạo lại",
		Code:    25,
	},
	{
		Key:     CommonKeyInvalidCarPartID,
		Message: "bộ phận xe không hợp lệ",
		Code:    26,
	},
	{
		Key:     CommonKeyInvalidUserID,
		Message: "người dùng không hợp lệ",
		Code:    27,
	},
	{
		Key:     CommonKeyInvalidKeyword,
		Message: "từ khóa không hợp lệ",
		Code:    28,
	},
	{
		Key:     CommonKeyInvalidNotificationID,
		Message: "id tin nhắn không hợp lệ",
		Code:    29,
	},
	{
		Key:     CommonKeErrorWhenUpdateData,
		Message: "đã có lỗi khi cập nhật dữ liệu, vui lòng thử lại",
		Code:    30,
	},
}
