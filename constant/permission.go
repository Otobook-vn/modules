package constant

import "strings"

// List permissions
const (
	StaffView  = "staff_view"
	StaffEdit  = "staff_edit"
	StaffOther = "staff_other"

	UserView  = "user_view"
	UserEdit  = "user_edit"
	UserOther = "user_other"

	ArticleView  = "article_view"
	ArticleEdit  = "article_edit"
	ArticleOther = "article_other"

	ConsultView  = "consult_view"
	ConsultEdit  = "consult_edit"
	ConsultOther = "consult_other"

	CarServiceView  = "car_service_view"
	CarServiceEdit  = "car_service_edit"
	CarServiceOther = "car_service_other"

	CarView  = "car_view"
	CarEdit  = "car_edit"
	CarOther = "car_other"
)

type list struct {
	Group       string       `json:"group"`
	Permissions []permission `json:"permissions"`
}

type permission struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func generatePermission(name, value string) permission {
	return permission{
		Name:  name,
		Value: strings.ToUpper(value),
	}
}

// ListStaffRoles ...
var ListStaffRoles = []list{
	{
		Group: "Quản lý nhân viên",
		Permissions: []permission{
			generatePermission("Xem thông tin", StaffView),
			generatePermission("Thêm/Chỉnh sửa", StaffEdit),
			generatePermission("Chức năng khác", StaffOther),
		},
	},
	{
		Group: "Quản lý khách hàng",
		Permissions: []permission{
			generatePermission("Xem thông tin", UserView),
			generatePermission("Thêm/Chỉnh sửa", UserEdit),
			generatePermission("Chức năng khác", UserOther),
		},
	},
	{
		Group: "Quản lý nội dung",
		Permissions: []permission{
			generatePermission("Xem thông tin", ArticleView),
			generatePermission("Thêm/Chỉnh sửa", ArticleEdit),
			generatePermission("Chức năng khác", ArticleOther),
		},
	},
	{
		Group: "Quản lý tư vấn",
		Permissions: []permission{
			generatePermission("Xem thông tin", ConsultView),
			generatePermission("Thêm/Chỉnh sửa", ConsultEdit),
			generatePermission("Chức năng khác", ConsultOther),
		},
	},
	{
		Group: "Quản lý dịch vụ",
		Permissions: []permission{
			generatePermission("Xem thông tin", CarServiceView),
			generatePermission("Thêm/Chỉnh sửa", CarServiceEdit),
			generatePermission("Chức năng khác", CarServiceOther),
		},
	},
	{
		Group: "Quản lý thông tin xe",
		Permissions: []permission{
			generatePermission("Xem thông tin", CarView),
			generatePermission("Thêm/Chỉnh sửa", CarEdit),
			generatePermission("Chức năng khác", CarOther),
		},
	},
}
