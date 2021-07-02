package filtervalue

import (
	"github.com/Otobook-vn/modules/constant"
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	"github.com/Otobook-vn/modules/model"
)

// UserListGroups ...
var UserListGroups = []filtermodel.CommonIDAndName{
	{
		ID: "all",
		Name: model.MultiLang{
			En: "All",
			Vi: "Tất cả",
		},
	},
	{
		ID: constant.UserGroupNormal,
		Name: model.MultiLang{
			En: "Normal",
			Vi: "Phổ thông",
		},
	},
	{
		ID: constant.UserGroupSpecialist,
		Name: model.MultiLang{
			En: "Specialist",
			Vi: "Chuyên gia",
		},
	},
}

// UserListRegisterFrom ...
var UserListRegisterFrom = []filtermodel.CommonIDAndName{
	{
		ID: constant.UserRegisterFromApp,
		Name: model.MultiLang{
			En: "App",
			Vi: "App",
		},
	},
}

// UserListStatuses ...
var UserListStatuses = []filtermodel.CommonIDAndName{
	{
		ID: constant.UserStatusActive,
		Name: model.MultiLang{
			En: "Active",
			Vi: "Hoạt động",
		},
	},
	{
		ID: constant.UserStatusBanned,
		Name: model.MultiLang{
			En: "Banned",
			Vi: "Khóa",
		},
	},
}
