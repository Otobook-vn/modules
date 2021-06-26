package filterget

import (
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	filtervalue "github.com/Otobook-vn/modules/filter/value"
	"github.com/Otobook-vn/modules/model"
	"github.com/thoas/go-funk"
)

// CommonIDAndNameValue ...
func CommonIDAndNameValue(list interface{}, id, lang string) model.ResponseCommonIDAndName {
	if itemInterface := funk.Find(list, func(i filtermodel.CommonIDAndName) bool {
		return i.ID == id
	}); itemInterface != nil {
		item := itemInterface.(filtermodel.CommonIDAndName)
		return model.ResponseCommonIDAndName{
			ID:   item.ID,
			Name: item.Name.GetValueByLang(lang),
		}
	}

	return model.ResponseCommonIDAndName{}
}

// Language ...
func Language(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.ListLanguages, id, lang)
}
