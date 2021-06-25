package constant

// Constant
var (
	UserGroupNormal     = "normal"
	UserGroupSpecialist = "specialist"
	UserGroupSystem     = "system"
	UserGroupList       = []interface{}{
		UserGroupNormal, UserGroupSpecialist, UserGroupSystem,
	}

	UserCodeOtobookVN = "otobookvn"

	UserStatusActive = "active"
	UserStatusBanned = "banned"

	UserRegisterFromApp = "app"
)
