package filtervalue

import (
	"github.com/Otobook-vn/modules/constant"
	filtermodel "github.com/Otobook-vn/modules/filter/model"
	"github.com/Otobook-vn/modules/model"
)

// NotificationManagementListCategories ...
var NotificationManagementListCategories = []filtermodel.CommonIDAndName{
	{
		ID: constant.NotificationCategoryCommon,
		Name: model.MultiLang{
			En: "Common",
			Vi: "Chung",
		},
	},
	{
		ID: constant.NotificationCategoryConsult,
		Name: model.MultiLang{
			En: "Consult",
			Vi: "Tư vấn",
		},
	},
}

// NotificationManagementListTargets ...
var NotificationManagementListTargets = []filtermodel.CommonIDAndName{
	{
		ID: constant.NotificationTargetAllUsers,
		Name: model.MultiLang{
			En: "All users",
			Vi: "Tất cả người dùng",
		},
	},
	{
		ID: constant.NotificationTargetByList,
		Name: model.MultiLang{
			En: "By list users",
			Vi: "Theo danh sách người dùng",
		},
	},
}

// NotificationManagementListStatuses ...
var NotificationManagementListStatuses = []filtermodel.CommonIDAndName{
	{
		ID: constant.NotificationSendStatusPending,
		Name: model.MultiLang{
			En: "Pending",
			Vi: "Chưa gửi",
		},
	},
	{
		ID: constant.NotificationSendStatusSent,
		Name: model.MultiLang{
			En: "Sent",
			Vi: "Đã gửi",
		},
	},
}

// NotificationManagementListActions ...
var NotificationManagementListActions = []filtermodel.NotificationManagementAction{
	{
		ID: constant.NotificationActionOpenAppPopup,
		Name: model.MultiLang{
			En: "Open app popup",
			Vi: "Mở app popup",
		},
		RequireValue: false,
	},
	{
		ID: constant.NotificationActionOpenExternalURL,
		Name: model.MultiLang{
			En: "Open external URL",
			Vi: "Mở liên kết ngoài",
		},
		RequireValue: true,
	},
	{
		ID: constant.NotificationActionOpenListCarServices,
		Name: model.MultiLang{
			En: "Open list Car services",
			Vi: "Mở danh sách Dịch vụ",
		},
		RequireValue: false,
	},
	{
		ID: constant.NotificationActionOpenCarServiceDetail,
		Name: model.MultiLang{
			En: "Open Car service detail",
			Vi: "Mở danh sách Dịch vụ",
		},
		RequireValue: true,
	},
	{
		ID: constant.NotificationActionOpenListArticles,
		Name: model.MultiLang{
			En: "Open list Articles",
			Vi: "Mở danh sách Bài viết",
		},
		RequireValue: false,
	},
	{
		ID: constant.NotificationActionOpenArticleDetail,
		Name: model.MultiLang{
			En: "Open Article detail",
			Vi: "Mở chi tiết Bài viết",
		},
		RequireValue: true,
	},
	{
		ID: constant.NotificationActionOpenListConsultQuestions,
		Name: model.MultiLang{
			En: "Open list Consult questions",
			Vi: "Mở danh sách Câu hỏi tư vấn",
		},
		RequireValue: false,
	},
	{
		ID: constant.NotificationActionOpenConsultQuestionDetail,
		Name: model.MultiLang{
			En: "Open consult question detail",
			Vi: "Mở chi tiết Câu hỏi tư vấn",
		},
		RequireValue: true,
	},
	{
		ID: constant.NotificationActionOpenListNotifications,
		Name: model.MultiLang{
			En: "Open list Notifications",
			Vi: "Mở danh sách Thông báo",
		},
		RequireValue: false,
	},
}
