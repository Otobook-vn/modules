package constant

// Constant
var (
	ArticleCategoryCodeInformation = "information"
	ArticleCategoryCodePromotion   = "promotion"
	ArticleCategoryCodeNeedToKnow  = "need_to_know"
	ArticleCategoryList            = []interface{}{
		ArticleCategoryCodeInformation, ArticleCategoryCodePromotion, ArticleCategoryCodeNeedToKnow,
	}

	ArticleDisplayTypeOpenURL    = "open_url"
	ArticleDisplayTypeSelfRender = "self_render"
	ArticleDisplayTypeList       = []interface{}{
		ArticleDisplayTypeOpenURL, ArticleDisplayTypeSelfRender,
	}

	ArticleStatusDraft     = "draft"
	ArticleStatusDisabled  = "disabled"
	ArticleStatusCompleted = "completed"
	ArticleStatusPublished = "published"
	ArticleStatusList      = []interface{}{
		ArticleStatusDraft, ArticleStatusDisabled, ArticleStatusCompleted, ArticleStatusPublished,
	}
)
