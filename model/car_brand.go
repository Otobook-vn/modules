package model

import (
	"strings"
	"time"

	"github.com/Otobook-vn/modules/utils"
)

// CarBrand ...
type CarBrand struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Slug         string
	Status       string
	IsOther      bool
	SearchTokens TsVector
	CreatedAt    time.Time

	// Ref
	Founded   Country `gorm:"foreignKey:FoundedID"`
	FoundedID *string
}

// TableName overrides the table name
func (CarBrand) TableName() string {
	return "car_brands"
}

// GenerateSearchTokens ...
func (s *CarBrand) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(s.Name)}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
