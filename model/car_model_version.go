package model

import (
	"strings"
	"time"

	"github.com/Otobook-vn/modules/utils"
)

// CarModelVersion ...
type CarModelVersion struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Slug         string
	NumOfSeat    int
	NumOfAirbag  int
	Status       string
	IsOther      bool
	SearchTokens TsVector
	CreatedAt    time.Time

	// Ref
	Brand          CarBrand `gorm:"foreignKey:BrandID"`
	BrandID        string
	Model          CarModel `gorm:"foreignKey:ModelID"`
	ModelID        string
	Transmission   CarTransmission `gorm:"foreignKey:TransmissionID"`
	TransmissionID *string
	BodyStyle      CarBodyStyle `gorm:"foreignKey:BodyStyleID"`
	BodyStyleID    *string
}

// TableName overrides the table name
func (CarModelVersion) TableName() string {
	return "car_model_versions"
}

// GenerateSearchTokens ...
func (s *CarModelVersion) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(s.Name)}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
