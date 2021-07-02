package filtervalue

import (
	"strconv"
	"time"
)

// CarListYears ...
var CarListYears = allYears()

// allYears ...
func allYears() []string {
	result := make([]string, 0)
	currentYear := time.Now().Year()
	for i := currentYear; i >= 1990; i-- {
		year := strconv.Itoa(i)
		result = append(result, year)
		if i == 1990 {
			result = append(result, "< "+year)
		}
	}
	return result
}
