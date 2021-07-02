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

// UserListRegisterFroms ...
var UserListRegisterFroms = []filtermodel.CommonIDAndName{
	{
		ID: constant.UserRegisterFromApp,
		Name: model.MultiLang{
			En: "App",
			Vi: "App",
		},
	},
}
