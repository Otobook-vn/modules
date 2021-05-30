package utils

import "github.com/google/uuid"

// GenerateUUID ...
func GenerateUUID() string {
	value, _ := uuid.NewRandom()
	return value.String()
}
