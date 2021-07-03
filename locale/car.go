package locale

// Constant
const (
	CarKeyBrandNameRequired              = "car_brand_name_required"
	CarKeyBrandNameExisted               = "car_brand_name_existed"
	CarKeyBrandNotFound                  = "car_brand_not_found"
	CarKeyBrandCannotEditOther           = "car_brand_cannot_edit_other"
	CarKeyModelNameRequired              = "car_model_name_required"
	CarKeyModelNameExisted               = "car_model_name_existed"
	CarKeyModelNotFound                  = "car_model_not_found"
	CarKeyModelCannotEditOther           = "car_model_cannot_edit_other"
	CarKeyModelVersionNameRequired       = "car_model_version_name_required"
	CarKeyModelVersionNameExisted        = "car_model_version_name_existed"
	CarKeyModelVersionNotFound           = "car_model_version_not_found"
	CarKeyModelVersionCannotEditOther    = "car_model_version_cannot_edit_other"
	CarKeyModelVersionInvalidNumOfSeat   = "car_model_version_invalid_num_of_seat"
	CarKeyModelVersionInvalidNumOfAirbag = "car_model_version_invalid_num_of_airbag"
	CarKeyInvalidTransmission            = "car_invalid_transmission"
	CarKeyInvalidBodyStyle               = "car_invalid_body_style"
	CarKeyPartInvalidName                = "car_part_invalid_name"
	CarKeyErrorCodeInvalidCode           = "car_error_code_invalid_code"
	CarKeyErrorCodeExisted               = "car_error_code_existed"
	CarKeyErrorCodeInvalidName           = "car_error_code_invalid_name"
	CarKeyErrorCodeInvalidDesc           = "car_error_code_invalid_desc"
	CarKeyPartNameExisted                = "car_part_name_existed"
)

// 300-349
var car = []Locale{
	{
		Key:     CarKeyBrandNameRequired,
		Message: "tên hãng xe không được trống",
		Code:    300,
	},
	{
		Key:     CarKeyBrandNameExisted,
		Message: "tên hãng xe đã tồn tại",
		Code:    301,
	},
	{
		Key:     CarKeyBrandNotFound,
		Message: "hãng xe không tìm thấy",
		Code:    302,
	},
	{
		Key:     CarKeyBrandCannotEditOther,
		Message: "không thể chỉnh sửa hãng Khác",
		Code:    303,
	},
	{
		Key:     CarKeyModelNameRequired,
		Message: "tên dòng xe không được trống",
		Code:    304,
	},
	{
		Key:     CarKeyModelNameExisted,
		Message: "tên dòng xe đã tồn tại",
		Code:    305,
	},
	{
		Key:     CarKeyModelNotFound,
		Message: "dòng xe không tìm thấy",
		Code:    306,
	},
	{
		Key:     CarKeyModelCannotEditOther,
		Message: "không thể chỉnh sửa dòng Khác",
		Code:    307,
	},
	{
		Key:     CarKeyModelVersionNameRequired,
		Message: "tên phiên bản không được trống",
		Code:    308,
	},
	{
		Key:     CarKeyModelVersionNotFound,
		Message: "phiên bản không tìm thấy",
		Code:    309,
	},
	{
		Key:     CarKeyModelVersionCannotEditOther,
		Message: "không thể chỉnh sửa phiên bản Khác",
		Code:    310,
	},
	{
		Key:     CarKeyModelVersionInvalidNumOfSeat,
		Message: "số ghế không hợp lệ",
		Code:    311,
	},
	{
		Key:     CarKeyModelVersionInvalidNumOfAirbag,
		Message: "số túi khí không hợp lệ",
		Code:    312,
	},
	{
		Key:     CarKeyInvalidTransmission,
		Message: "hộp số không hợp lệ",
		Code:    313,
	},
	{
		Key:     CarKeyInvalidBodyStyle,
		Message: "kiểu dáng không hợp lệ",
		Code:    314,
	},
	{
		Key:     CarKeyModelVersionNameExisted,
		Message: "tên phiên bản đã tồn tại",
		Code:    315,
	},
	{
		Key:     CarKeyPartInvalidName,
		Message: "tên bộ phận không hợp lệ",
		Code:    316,
	},
	{
		Key:     CarKeyErrorCodeInvalidCode,
		Message: "mã lỗi không hợp lệ",
		Code:    317,
	},
	{
		Key:     CarKeyErrorCodeExisted,
		Message: "mã lỗi đã tồn tại với hãng xe này",
		Code:    318,
	},
	{
		Key:     CarKeyErrorCodeInvalidName,
		Message: "tên mã lỗi không hợp lệ",
		Code:    319,
	},
	{
		Key:     CarKeyErrorCodeInvalidDesc,
		Message: "mô tả mã lỗi không hợp lệ",
		Code:    320,
	},
	{
		Key:     CarKeyPartNameExisted,
		Message: "tên bộ phận đã tồn tại",
		Code:    321,
	},
}
