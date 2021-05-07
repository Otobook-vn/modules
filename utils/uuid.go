package utils

import "github.com/google/uuid"

func NewUUID() string {
	value, _ := uuid.NewRandom()
	return value.String()
}
