package constant

// Constant
var (
	NotificationPushYes = true
	NotificationPushNo  = false

	NotificationCategoryCommon  = "common"
	NotificationCategoryConsult = "consult"
	NotificationCategoryList    = []interface{}{
		NotificationCategoryCommon, NotificationCategoryConsult,
	}

	NotificationSendOptionAllUser = "all_users"
	NotificationSendOptionByList  = "by_list"
	NotificationSendOptionList    = []interface{}{
		NotificationSendOptionAllUser, NotificationSendOptionByList,
	}

	NotificationTypeConsultQuestionAdminChangeStatus       = "consult_question_admin_change_status"
	NotificationTypeConsultQuestionSpecialistMarkCompleted = "consult_question_specialist_mark_completed"
	NotificationTypeConsultQuestionHasNewComment           = "consult_question_has_new_comment"
	NotificationTypeConsultQuestionUserRating              = "consult_question_user_rating"
	NotificationTypeConsultQuestionResolved                = "consult_question_resolved"

	NotificationActionOpenConsultQuestionDetail = "open_consult_question_detail"
)
