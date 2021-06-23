package translation

import (
	"fmt"

	"github.com/Otobook-vn/modules/constant"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

var listNotificationContent = make([]languageData, 0)

func loadNotificationContent(bundle *i18n.Bundle, pwd string) {
	for _, langInterface := range constant.LangList {
		lang := langInterface.(string)

		// Add language
		path := fmt.Sprintf("%s/notification/content/active.%s.toml", pwd, lang)
		_, err := bundle.LoadMessageFile(path)
		if err != nil {
			panic(err)
		}

		// Add localizer
		localizer := i18n.NewLocalizer(bundle, lang)

		// Add item
		item := languageData{
			Lang:      lang,
			Localizer: localizer,
		}
		listNotificationContent = append(listNotificationContent, item)
	}
}

// GetNotificationContent ...
func GetNotificationContent(lang, messageID string, data interface{}) string {
	itemInterface := funk.Find(listNotificationContent, func(i languageData) bool {
		return i.Lang == lang
	})
	if itemInterface == nil {
		return ""
	}

	item := itemInterface.(languageData)
	msg, _ := item.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	return msg
}
