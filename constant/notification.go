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
	NotificationSendStatusSent    = "sent"
	NotificationSendStatusList    = []interface{}{
		NotificationSendStatusPending, NotificationSendStatusSent,
	}

	NotificationTypeConsultQuestionAdminChangeStatus       = "consult_question_admin_change_status"
	NotificationTypeConsultQuestionSpecialistMarkCompleted = "consult_question_specialist_mark_completed"
	NotificationTypeConsultQuestionHasNewComment           = "consult_question_has_new_comment"
	NotificationTypeConsultQuestionUserRating              = "consult_question_user_rating"
	NotificationTypeConsultQuestionResolved                = "consult_question_resolved"
	NotificationTypeConsultQuestionUserLiked               = "consult_question_user_liked"

	NotificationActionOpenAppPopup              = "open_app_popup"
	NotificationActionOpenExternalURL           = "open_external_url"
	NotificationActionOpenListCarServices       = "open_list_car_services"
	NotificationActionOpenCarServiceDetail      = "open_car_service_detail"
	NotificationActionOpenListArticles          = "open_list_articles"
	NotificationActionOpenArticleDetail         = "open_article_detail"
	NotificationActionOpenListConsultQuestions  = "open_list_consult_questions"
	NotificationActionOpenConsultQuestionDetail = "open_consult_question_detail"
	NotificationActionOpenListNotifications     = "open_list_notifications"
)
