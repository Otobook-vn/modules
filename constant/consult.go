package constant

// Constant
var (
	ConsultQuestionStatusReview    = "review"
	ConsultQuestionStatusOpen      = "open"
	ConsultQuestionStatusReject    = "reject"
	ConsultQuestionStatusCompleted = "completed"
	ConsultQuestionStatusResolved  = "resolved"
	ConsultQuestionStatusClosed    = "closed"
	ConsultQuestionStatusList      = []interface{}{
		ConsultQuestionStatusReview, ConsultQuestionStatusOpen, ConsultQuestionStatusReject,
		ConsultQuestionStatusCompleted, ConsultQuestionStatusResolved, ConsultQuestionStatusClosed,
	}

	ConsultHistoryActionUserCreateQuestion      = "user_create_question"
	ConsultHistoryActionUserEditQuestion        = "user_edit_question"
	ConsultHistoryActionUserReopenQuestion      = "user_reopen_question"
	ConsultHistoryActionUserDeleteQuestion      = "user_delete_question"
	ConsultHistoryActionAdminChangeStatus       = "admin_change_status"
	ConsultHistoryActionSpecialistMarkCompleted = "specialist_mark_completed"
	ConsultHistoryActionSystemAutoResolve       = "system_auto_resolve"
)
