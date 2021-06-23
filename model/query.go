package model

import (
	"time"

	"github.com/Otobook-vn/modules/utils"
)

// AppQuery ...
type AppQuery struct {
	Page            int
	Timestamp       time.Time
	Limit           int
	Keyword         string
	Status          string
	Role            string
	Group           string
	CarBrand        string
	CarModel        string
	CarTransmission string
	CarBodyStyle    string
	CarPart         string
	Province        int
	District        int
	Ward            int
	Category        string
	Latitude        float64
	Longitude       float64
	Sort            string
	Language        string
	Topic           string
	DateFrom        time.Time
	DateTo          time.Time
	User            string
	Code            string
}

// SetLimitMaxValue ...
func (q *AppQuery) SetLimitMaxValue() {
	if q.Limit > 50 || q.Limit < 0 {
		q.Limit = 50
	}
}

// IsInvalidLatLng ...
func (q AppQuery) IsInvalidLatLng() bool {
	return utils.IsInvalidLatLng(q.Latitude, q.Longitude)
}
