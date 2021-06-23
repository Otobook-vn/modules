package locale

// Constant
const (
	ConsultKeyCategoryInvalidName        = "consult_category_invalid_name"
	ConsultKeyCategoryNameExisted        = "consult_category_name_existed"
	ConsultKeyQuestionInvalidTitle       = "consult_question_invalid_title"
	ConsultKeyQuestionInvalidContent     = "consult_question_invalid_content"
	ConsultKeyQuestionInvalidCategory    = "consult_question_invalid_category"
	ConsultKeyQuestionNotFound           = "consult_question_not_found"
	ConsultKeyQuestionCannotEdit         = "consult_question_cannot_edit"
	ConsultKeyQuestionCannotComplete     = "consult_question_cannot_complete"
	ConsultKeyCommentInvalidContent      = "consult_comment_invalid_content"
	ConsultKeyQuestionCannotComment      = "consult_question_cannot_comment"
	ConsultKeyQuestionInvalidRatingPoint = "consult_question_invalid_rating_point"
	ConsultKeyQuestionCannotRating       = "consult_question_cannot_rating"
	ConsultKeyQuestionInvalidReason      = "consult_question_invalid_reason"
	ConsultKeyCommentInvalidQuestion     = "consult_comment_invalid_question"
)

// 450-499
var consult = []Locale{
	{
		Key:     ConsultKeyCategoryInvalidName,
		Message: "tên danh mục không hợp lệ",
		Code:    450,
	},
	{
		Key:     ConsultKeyCategoryNameExisted,
		Message: "tên danh mục đã tồn tại",
		Code:    451,
	},
	{
		Key:     ConsultKeyQuestionInvalidTitle,
		Message: "tiêu đề phải ít nhất 10 ký tự",
		Code:    452,
	},
	{
		Key:     ConsultKeyQuestionInvalidContent,
		Message: "nội dung phải ít nhất 30 ký tự",
		Code:    453,
	},
	{
		Key:     ConsultKeyQuestionInvalidCategory,
		Message: "danh mục câu hỏi không hợp lệ",
		Code:    454,
	},
	{
		Key:     ConsultKeyQuestionNotFound,
		Message: "câu hỏi không tìm thấy",
		Code:    455,
	},
	{
		Key:     ConsultKeyQuestionCannotEdit,
		Message: "câu hỏi hiện không thể chỉnh sửa",
		Code:    456,
	},
	{
		Key:     ConsultKeyQuestionCannotComplete,
		Message: "câu hỏi hiện không thể hoàn thành",
		Code:    457,
	},
	{
		Key:     ConsultKeyCommentInvalidContent,
		Message: "nội dung không được trống",
		Code:    458,
	},
	{
		Key:     ConsultKeyQuestionCannotComment,
		Message: "bạn không thể bình luận vào câu hỏi của người khác",
		Code:    459,
	},
	{
		Key:     ConsultKeyQuestionInvalidRatingPoint,
		Message: "điểm đánh giá không hợp lệ",
		Code:    460,
	},
	{
		Key:     ConsultKeyQuestionCannotRating,
		Message: "câu hỏi hiện không thể đánh giá",
		Code:    461,
	},
	{
		Key:     ConsultKeyQuestionInvalidReason,
		Message: "lý do không hợp lệ",
		Code:    462,
	},
	{
		Key:     ConsultKeyCommentInvalidQuestion,
		Message: "lý do không hợp lệ",
		Code:    463,
	},
}
