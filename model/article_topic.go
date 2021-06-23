package model

import (
	"strings"
	"time"

	"gorm.io/datatypes"
)

// ArticleTopic ...
type ArticleTopic struct {
	ID           string `gorm:"primaryKey"`
	Name         datatypes.JSON
	Desc         datatypes.JSON
	Status       string
	SearchTokens TsVector
	CreatedAt    time.Time

	// Ref
	Category     ArticleCategory `gorm:"foreignKey:CategoryCode"`
	CategoryCode string
}

// TableName overrides the table name
func (ArticleTopic) TableName() string {
	return "article_topics"
}

// GenerateSearchTokens ...
func (a *ArticleTopic) GenerateSearchTokens() {
	values := GetMultiLangFromJSON(a.Name).GetArrayValues()
	vecValue := strings.Join(values, " ")
	a.SearchTokens = TsVector{Value: vecValue}
}
