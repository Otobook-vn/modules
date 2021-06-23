package audit

// Sources
const (
	SourceOtobook = "otobook"
)

// List actions
const (
	ActionCreate            = "create"
	ActionUpdate            = "update"
	ActionUpdatePermissions = "update-permissions"
)

// List targets
const (
	TargetArticleTopic       = "article-topics"
	TargetArticle            = "articles"
	TargetCarServiceCategory = "car-service-categories"
	TargetCarService         = "car-services"
	TargetCarBrand           = "car-brands"
	TargetCarModel           = "car-models"
	TargetCarModelVersion    = "car-model-versions"
	TargetCarPart            = "car-parts"
	TargetCarErrorCode       = "car-error-codes"
	TargetStaff              = "staffs"
	TargetStaffRole          = "staff-roles"
)
