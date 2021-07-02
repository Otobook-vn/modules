package filterget

import (
	filtervalue "github.com/Otobook-vn/modules/filter/value"
	"github.com/Otobook-vn/modules/model"
)

// UserGroup ...
func UserGroup(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.UserListGroups, id, lang)
}

// UserRegisterFrom ...
func UserRegisterFrom(id, lang string) model.ResponseCommonIDAndName {
	return CommonIDAndNameValue(filtervalue.UserListRegisterFroms, id, lang)
}
