package filterget

import (
	filtervalue "github.com/Otobook-vn/modules/filter/value"
	"github.com/Otobook-vn/modules/model"
)

// ArticleCategory ...
func ArticleCategory(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.ArticleCategories, id, lang)
}
