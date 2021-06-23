package model

import (
	"strings"
	"time"

	"github.com/Otobook-vn/modules/utils"
)

// CarModel ...
type CarModel struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Slug         string
	Status       string
	IsOther      bool
	SearchTokens TsVector
	CreatedAt    time.Time

	// Ref
	Brand   CarBrand `gorm:"foreignKey:BrandID"`
	BrandID string
}

// TableName overrides the table name
func (CarModel) TableName() string {
	return "car_models"
}

// GenerateSearchTokens ...
func (s *CarModel) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(s.Name)}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
