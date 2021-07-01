package model

import (
	"strings"
	"time"

	"github.com/Otobook-vn/modules/utils"
)

// User ...
type User struct {
	ID               string `gorm:"primaryKey"`
	Group            string
	Code             string
	Name             string
	PhoneCountryCode string
	PhoneNumber      string
	PhoneFull        string
	Email            string
	Avatar           string
	HashedPassword   string
	Status           string
	ReferralCode     string
	Desc             string
	LocationAddress  string
	SearchTokens     TsVector
	IsPhoneVerified  bool
	RegisterFrom     string
	PhoneVerifiedAt  time.Time
	LastActivatedAt  time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// Ref
	LocationProvince   LocationProvince `gorm:"foreignKey:LocationProvinceID"`
	LocationProvinceID *int
	LocationDistrict   LocationDistrict `gorm:"foreignKey:LocationDistrictID"`
	LocationDistrictID *int
	LocationWard       LocationWard `gorm:"foreignKey:LocationWardID"`
	LocationWardID     *int
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

// GenerateSearchTokens ...
func (u *User) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(u.Name), u.PhoneFull, u.Email}
	vecValue := strings.Join(values, " ")
	u.SearchTokens = TsVector{Value: vecValue}
}
