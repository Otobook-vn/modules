package filterget

import (
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	filtervalue "github.com/Otobook-vn/modules/filter/value"
	"github.com/Otobook-vn/modules/model"
	"github.com/thoas/go-funk"
)

// NotificationManagementAction ...
func NotificationManagementAction(id, lang string) model.ResponseCommonIDAndName {
	if itemInterface := funk.Find(filtervalue.NotificationManagementListActions, func(i filtermodel.NotificationManagementAction) bool {
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

// NotificationManagementCategory ...
func NotificationManagementCategory(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.NotificationManagementListCategories, id, lang)
}

// NotificationManagementTarget ...
func NotificationManagementTarget(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.NotificationManagementListTargets, id, lang)
}

// NotificationManagementStatus ...
func NotificationManagementStatus(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.NotificationManagementListStatuses, id, lang)
}
