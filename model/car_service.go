package model

import (
	"strings"
	"time"

	"github.com/Otobook-vn/modules/utils"
	"gorm.io/datatypes"
)

// CarService ...
type CarService struct {
	ID            string `gorm:"primaryKey"`
	Name          datatypes.JSON
	Phone         string
	Email         string
	TaxIDNumber   string
	Desc          datatypes.JSON
	Status        string
	Rating        float64
	Address       string
	SearchTokens  TsVector
	IsVerified    bool
	LocationPoint LocationPoint
	ContactName   string
	ContactEmail  string
	ContactPhone  string
	Photos        datatypes.JSON
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Ref
	Categories         []CarServiceCategory `gorm:"many2many:car_services_and_categories"`
	Country            Country              `gorm:"foreignKey:CountryCode"`
	CountryCode        string
	LocationProvince   LocationProvince `gorm:"foreignKey:LocationProvinceID"`
	LocationProvinceID *int
	LocationDistrict   LocationDistrict `gorm:"foreignKey:LocationDistrictID"`
	LocationDistrictID *int
	LocationWard       LocationWard `gorm:"foreignKey:LocationWardID"`
	LocationWardID     *int
}

// TableName overrides the table name
func (CarService) TableName() string {
	return "car_services"
}

// GenerateSearchTokens ...
func (s *CarService) GenerateSearchTokens() {
	nameMultiLang := GetMultiLangFromJSON(s.Name)
	values := []string{utils.RemoveDiacritics(nameMultiLang.Vi), utils.RemoveDiacritics(nameMultiLang.En), s.Phone, s.Email, s.TaxIDNumber}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
