package filtervalue

import (
	"github.com/Otobook-vn/modules/constant"
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	"github.com/Otobook-vn/modules/model"
)

// ArticleCategories ...
var ArticleCategories = []filtermodel.CommonIDAndName{
	{
		ID: constant.ArticleCategoryCodeInformation,
		Name: model.MultiLang{
			En: "Information",
			Vi: "Thông tin",
		},
	},
	{
		ID: constant.ArticleCategoryCodePromotion,
		Name: model.MultiLang{
			En: "Promotion",
			Vi: "Khuyến mãi",
		},
	},
	{
		ID: constant.ArticleCategoryCodeNeedToKnow,
		Name: model.MultiLang{
			En: "Need to know",
			Vi: "Cần biết",
		},
	},
}
