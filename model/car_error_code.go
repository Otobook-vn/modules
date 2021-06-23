package model

import (
	"strings"
	"time"

	"gorm.io/datatypes"
)

// CarErrorCode ...
type CarErrorCode struct {
	ID           string `gorm:"primaryKey"`
	Code         string
	Name         datatypes.JSON
	Desc         datatypes.JSON
	Status       string
	SearchTokens TsVector
	CreatedAt    time.Time

	// Ref
	Brand   CarBrand `gorm:"foreignKey:BrandID"`
	BrandID string
	Part    CarPart `gorm:"foreignKey:PartID"`
	PartID  string
}

// TableName overrides the table name
func (CarErrorCode) TableName() string {
	return "car_error_codes"
}

// GenerateSearchTokens ...
func (s *CarErrorCode) GenerateSearchTokens() {
	nameValues := GetMultiLangFromJSON(s.Name).GetArrayValues()
	descValues := GetMultiLangFromJSON(s.Desc).GetArrayValues()
	values := append(nameValues, descValues...)
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
