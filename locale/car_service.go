package locale

// Constant ...
const (
	CarServiceKeyNameRequired         = "car_service_name_required"
	CarServiceKeyPhoneRequired        = "car_service_phone_required"
	CarServiceCategoryKeyNameRequired = "car_service_category_name_required"
	CarServiceCategoryKeyNameExisted  = "car_service_category_name_existed"
)

// 350-399
var service = []Locale{
	{
		Key:     CarServiceKeyNameRequired,
		Message: "tên không được trống",
		Code:    350,
	},
	{
		Key:     CarServiceKeyPhoneRequired,
		Message: "số điện thoại không được trống",
		Code:    351,
	},
	{
		Key:     CarServiceCategoryKeyNameRequired,
		Message: "tên danh mục không được trống",
		Code:    352,
	},
	{
		Key:     CarServiceCategoryKeyNameExisted,
		Message: "tên danh mục đã tồn tại",
		Code:    353,
	},
}
