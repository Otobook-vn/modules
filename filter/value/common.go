package filtervalue

import (
	"github.com/Otobook-vn/modules/constant"
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	"github.com/Otobook-vn/modules/model"
)

// ListLanguages ...
var ListLanguages = []filtermodel.CommonIDAndName{
	{
		ID: constant.LangVi,
		Name: model.MultiLang{
			En: "Vietnamese",
			Vi: "Tiếng Việt",
		},
	},
	{
		ID: constant.LangEn,
		Name: model.MultiLang{
			En: "English",
			Vi: "Tiếng Anh",
		},
	},
}
