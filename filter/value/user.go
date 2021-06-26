package filtervalue

import (
	"github.com/Otobook-vn/modules/constant"
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	"github.com/Otobook-vn/modules/model"
)

// UserListGroups ...
var UserListGroups = []filtermodel.CommonIDAndName{
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
