package model

import (
	"strings"
	"time"

	"gorm.io/datatypes"
)

// ArticleCategory ...
type ArticleCategory struct {
	ID           string
	Code         string `gorm:"primaryKey"`
	Name         datatypes.JSON
	Status       string
	SearchTokens TsVector
	CreatedAt    time.Time
}

// TableName overrides the table name
func (ArticleCategory) TableName() string {
	return "article_categories"
}

// GenerateSearchTokens ...
func (a *ArticleCategory) GenerateSearchTokens() {
	values := GetMultiLangFromJSON(a.Name).GetArrayValues()
	vecValue := strings.Join(values, " ")
	a.SearchTokens = TsVector{Value: vecValue}
}
