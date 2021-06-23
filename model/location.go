package model

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/twpayne/go-geom/encoding/wkbhex"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Otobook-vn/modules/utils"
)

//
// LOCATION POINT
//

// LocationPoint ...
type LocationPoint struct {
	Latitude, Longitude float64
}

// GormDataType ...
func (loc LocationPoint) GormDataType() string {
	return "geometry"
}

// String ...
func (loc LocationPoint) String() string {
	return fmt.Sprintf("POINT(%f %f)", loc.Latitude, loc.Longitude)
}

// GormValue ...
func (loc LocationPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.Latitude, loc.Longitude)},
	}
}

// Scan implements the sql.Scanner interface
func (loc *LocationPoint) Scan(v interface{}) error {
	valueString, ok := v.(string)
	if !ok {
		return nil
	}
	data, err := wkbhex.Decode(valueString)
	if err != nil {
		return nil
	}

	// Assign value
	flatCoords := data.FlatCoords()
	loc.Latitude = flatCoords[0]
	loc.Longitude = flatCoords[1]

	return nil
}

// IsInvalidLatLng ...
func (loc LocationPoint) IsInvalidLatLng() bool {
	return utils.IsInvalidLatLng(loc.Latitude, loc.Longitude)
}

// CalculateDistance ...
func (loc LocationPoint) CalculateDistance(lat, lon float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = loc.Latitude * math.Pi / 180
	lo1 = loc.Longitude * math.Pi / 180
	la2 = lat * math.Pi / 180
	lo2 = lon * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := utils.Hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*utils.Hsin(lo2-lo1)

	dist := 2 * r * math.Asin(math.Sqrt(h))
	return math.Round(dist)
}

//
// LOCATION PROVINCE
//

// LocationProvince ...
type LocationProvince struct {
	ID           string
	Name         string
	SearchTokens TsVector
	Code         int `gorm:"primaryKey"`
	CreatedAt    time.Time
}

// TableName overrides the table name
func (LocationProvince) TableName() string {
	return "location_provinces"
}

// GenerateSearchTokens ...
func (l *LocationProvince) GenerateSearchTokens() {
	l.SearchTokens = locationGenerateSearchToken(l.Name, l.Code)
}

//
// LOCATION DISTRICT
//

// LocationDistrict ...
type LocationDistrict struct {
	ID           string
	Name         string
	SearchTokens TsVector
	Code         int `gorm:"primaryKey"`
	CreatedAt    time.Time

	// Ref
	Province   LocationProvince `gorm:"foreignKey:ProvinceID"`
	ProvinceID int
}

// TableName overrides the table name
func (LocationDistrict) TableName() string {
	return "location_districts"
}

// GenerateSearchTokens ...
func (l *LocationDistrict) GenerateSearchTokens() {
	l.SearchTokens = locationGenerateSearchToken(l.Name, l.Code)
}

//
// LOCATION WARD
//

// LocationWard ...
type LocationWard struct {
	ID           string
	Name         string
	SearchTokens TsVector
	Code         int `gorm:"primaryKey"`
	CreatedAt    time.Time

	// Ref
	Province   LocationProvince `gorm:"foreignKey:ProvinceID"`
	ProvinceID int
	District   LocationDistrict `gorm:"foreignKey:DistrictID"`
	DistrictID int
}

// GenerateSearchTokens ...
func (l *LocationWard) GenerateSearchTokens() {
	l.SearchTokens = locationGenerateSearchToken(l.Name, l.Code)
}

// TableName overrides the table name
func (LocationWard) TableName() string {
	return "location_wards"
}

func locationGenerateSearchToken(name string, code int) TsVector {
	codeString := strconv.Itoa(code)
	values := []string{utils.RemoveDiacritics(name) + codeString}
	vecValue := strings.Join(values, " ")
	return TsVector{Value: vecValue}
}
