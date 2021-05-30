package utils

import (
	"encoding/json"

	"gorm.io/datatypes"
)

// ConvertToDataTypesJSON ...
func ConvertToDataTypesJSON(d interface{}) datatypes.JSON {
	bytes, _ := json.Marshal(d)

	// JSON data
	jsonData := datatypes.JSON(bytes)
	return jsonData
}
