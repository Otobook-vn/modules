package model

import (
	"strings"
	"time"

	"gorm.io/datatypes"
)

// ConsultCategory ...
type ConsultCategory struct {
	ID           string `gorm:"primaryKey"`
	Name         datatypes.JSON
	SearchTokens TsVector
	Slug         string
	Status       string
	CreatedAt    time.Time
}

// TableName overrides the table name
func (ConsultCategory) TableName() string {
	return "consult_categories"
}

// GenerateSearchTokens ...
func (s *ConsultCategory) GenerateSearchTokens() {
	values := GetMultiLangFromJSON(s.Name).GetArrayValues()
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
