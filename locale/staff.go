package locale

// Constant
const (
	StaffKeyUsernameExisted         = "staff_username_existed"
	StaffKeyUsernameRequired        = "staff_username_required"
	StaffKeyUsernameInvalidLength   = "staff_username_invalid_length"
	StaffKeyNameRequired            = "staff_name_required"
	StaffKeyPasswordRequired        = "staff_password_required"
	StaffKeyPasswordInvalidLength   = "staff_password_invalid_length"
	StaffKeyRoleRequired            = "staff_role_required"
	StaffKeyStatusRequired          = "staff_status_required"
	StaffRoleKeyInvalidName         = "staff_role_invalid_name"
	StaffRoleKeyNameExisted         = "staff_role_name_existed"
	StaffRoleKeyAdminCannotBeEdited = "staff_role_admin_cannot_be_edited"
)

// 200-249
var staff = []Locale{
	{
		Key:     StaffKeyUsernameExisted,
		Message: "username đã tồn tại",
		Code:    201,
	},
	{
		Key:     StaffKeyUsernameRequired,
		Message: "username không được trống",
		Code:    202,
	},
	{
		Key:     StaffKeyNameRequired,
		Message: "tên không được trống",
		Code:    203,
	},
	{
		Key:     StaffKeyPasswordRequired,
		Message: "mật khẩu không được trống",
		Code:    204,
	}, {
		Key:     StaffKeyPasswordInvalidLength,
		Message: "mật khẩu phải từ 6-32 ký tự",
		Code:    205,
	},
	{
		Key:     StaffKeyRoleRequired,
		Message: "nhóm quyền không được trống",
		Code:    206,
	}, {
		Key:     StaffKeyUsernameInvalidLength,
		Message: "username phải từ 3-32 ký tự",
		Code:    207,
	},
	{
		Key:     StaffKeyStatusRequired,
		Message: "trạng thái không được trống",
		Code:    208,
	},
	{
		Key:     StaffRoleKeyInvalidName,
		Message: "tên nhóm quyền không hợp lệ",
		Code:    209,
	},
	{
		Key:     StaffRoleKeyNameExisted,
		Message: "tên nhóm quyền đã tồn tại",
		Code:    210,
	},
	{
		Key:     StaffRoleKeyAdminCannotBeEdited,
		Message: "không thể chỉnh sửa nhóm Admin",
		Code:    211,
	},
}
