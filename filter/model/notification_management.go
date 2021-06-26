package filtermodel

import "github.com/Otobook-vn/modules/model"

type (
	// NotificationManagementAction ...
	NotificationManagementAction struct {
		ID           string          `json:"id"`
		Name         model.MultiLang `json:"name"`
		RequireValue bool            `json:"requireValue"`
	}
)
