package locale

// Constant
const (
	ArticleKeyTitleRequired          = "article_title_required"
	ArticleKeyInvalidCategory        = "article_invalid_category"
	ArticleKeyInvalidTopic           = "article_invalid_topic"
	ArticleKeyDisplayURLRequired     = "article_display_url_required"
	ArticleKeyDisplayContentRequired = "article_display_content_required"
	ArticleKeyTopicInvalidName       = "article_topic_invalid_name"
	ArticleKeyTopicInvalidCategory   = "article_topic_invalid_category"
	ArticleKeyInvalidDisplayType     = "article_invalid_display_type"
	ArticleKeyTopicExisted           = "article_topic_existed"
)

// 400-449
var article = []Locale{
	{
		Key:     ArticleKeyTitleRequired,
		Message: "tiêu đề bài viết không được trống",
		Code:    400,
	},
	{
		Key:     ArticleKeyInvalidCategory,
		Message: "danh mục bài viết không hợp lệ",
		Code:    401,
	},
	{
		Key:     ArticleKeyInvalidTopic,
		Message: "chủ đề bài viết không hợp lệ",
		Code:    402,
	},
	{
		Key:     ArticleKeyDisplayURLRequired,
		Message: "đường dẫn bài viết không được trống",
		Code:    403,
	},
	{
		Key:     ArticleKeyDisplayContentRequired,
		Message: "nội dung bài viết không được trống",
		Code:    404,
	},
	{
		Key:     ArticleKeyTopicInvalidName,
		Message: "tên chủ đề không được trống",
		Code:    405,
	},
	{
		Key:     ArticleKeyTopicInvalidCategory,
		Message: "danh mục chủ đề không hợp lệ",
		Code:    406,
	},
	{
		Key:     ArticleKeyInvalidDisplayType,
		Message: "loại hiển thị không hợp lệ",
		Code:    407,
	},
	{
		Key:     ArticleKeyTopicExisted,
		Message: "tên chủ đề đã tồn tại",
		Code:    408,
	},
}
