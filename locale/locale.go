package locale

import (
	"github.com/thoas/go-funk"
)

// Locale ...
type Locale struct {
	Key     string
	Message string
	Code    int
}

var notFoundKey = Locale{
	Key:     CommonKeyNotFound,
	Message: "không tìm thấy",
	Code:    -1,
}

// GetByKey give key and receive message + code
func GetByKey(key string) Locale {
	item := funk.Find(list, func(item Locale) bool {
		return item.Key == key
	})

	if item == nil {
		return notFoundKey
	}
	return item.(Locale)
}

var list = make([]Locale, 0)

// Init ...
func Init() {
	// 1-99
	list = append(list, common...)
	// 200-249
	list = append(list, staff...)
	// 250-299
	list = append(list, user...)
	// 300-349
	list = append(list, car...)
	// 350-399
	list = append(list, service...)
	// 400-449
	list = append(list, article...)
	// 450-499
	list = append(list, consult...)
}
