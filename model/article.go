package model

import (
	"strings"
	"time"

	"gorm.io/datatypes"

	"github.com/Otobook-vn/modules/utils"
)

// Article ...
type Article struct {
	ID             string `gorm:"primaryKey"`
	Title          string
	Thumbnail      datatypes.JSON
	Desc           string
	DisplayType    string // url: open url in web view, self: app render with content
	DisplayURL     string // other source url
	DisplayContent string // self render
	Status         string
	Language       string
	StatsView      int
	SearchTokens   TsVector
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// Ref
	Category     ArticleCategory `gorm:"foreignKey:CategoryCode"`
	CategoryCode string
	Topics       []ArticleTopic `gorm:"many2many:articles_and_topics"`
}

// TableName overrides the table name
func (Article) TableName() string {
	return "articles"
}

// GenerateSearchTokens ...
func (a *Article) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(a.Title)}
	vecValue := strings.Join(values, " ")
	a.SearchTokens = TsVector{Value: vecValue}
}
