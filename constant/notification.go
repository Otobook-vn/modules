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

	NotificationTargetAllUsers = "all_users"
	NotificationTargetByList   = "by_list"
	NotificationTargetList     = []interface{}{
		NotificationTargetAllUsers, NotificationTargetByList,
	}

	NotificationSendStatusPending = "pending"
	NotificationSendStatusSent    = "send"
	NotificationSendStatusList    = []interface{}{
		NotificationSendStatusPending, NotificationSendStatusSent,
	}

	NotificationTypeConsultQuestionAdminChangeStatus       = "consult_question_admin_change_status"
	NotificationTypeConsultQuestionSpecialistMarkCompleted = "consult_question_specialist_mark_completed"
	NotificationTypeConsultQuestionHasNewComment           = "consult_question_has_new_comment"
	NotificationTypeConsultQuestionUserRating              = "consult_question_user_rating"
	NotificationTypeConsultQuestionResolved                = "consult_question_resolved"

	NotificationActionOpenConsultQuestionDetail = "open_consult_question_detail"
)
