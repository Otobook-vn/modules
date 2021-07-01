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
	ActionAddUsers          = "add-users"
	ActionRemoveUsers       = "remove-users"
	ActionSend              = "send"
)

// List targets
const (
	TargetOtobookArticleTopic           = "article-topics"
	TargetOtobookArticle                = "articles"
	TargetOtobookCarServiceCategory     = "car-service-categories"
	TargetOtobookCarService             = "car-services"
	TargetOtobookCarBrand               = "car-brands"
	TargetOtobookCarModel               = "car-models"
	TargetOtobookCarModelVersion        = "car-model-versions"
	TargetOtobookCarPart                = "car-parts"
	TargetOtobookCarErrorCode           = "car-error-codes"
	TargetOtobookStaff                  = "staffs"
	TargetOtobookStaffRole              = "staff-roles"
	TargetOtobookConsultCategory        = "consult-categories"
	TargetOtobookConsultComment         = "consult-comments"
	TargetOtobookNotificationManagement = "notification-managements"
	TargetOtobookUser                   = "users"
)

// OtobookTargets ...
var OtobookTargets = []string{
	TargetOtobookArticleTopic,
	TargetOtobookArticle,
	TargetOtobookCarServiceCategory,
	TargetOtobookCarService,
	TargetOtobookCarBrand,
	TargetOtobookCarModel,
	TargetOtobookCarModelVersion,
	TargetOtobookCarPart,
	TargetOtobookCarErrorCode,
	TargetOtobookStaff,
	TargetOtobookStaffRole,
	TargetOtobookConsultCategory,
	TargetOtobookConsultComment,
	TargetOtobookNotificationManagement,
}
