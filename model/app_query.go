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
	Email           string
	Name            string
	Phone           string
	Status          string
	Role            string
	Group           string
	CarBrand        string
	CarModel        string
	CarVersion      string
	CarTransmission string
	CarBodyStyle    string
	CarPart         string
	CarColor        string
	CarLicensePlate string
	CarVINNumber    string
	Year            string
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
	Target          string
	Sender          string
	Notification    string
	RegisterFrom    string
	Exclude         string
}

// SetLimitMaxValue ...
func (q *AppQuery) SetLimitMaxValue() {
	if q.Limit > 50 || q.Limit < 1 {
		q.Limit = 50
	}
}

// IsInvalidLatLng ...
func (q AppQuery) IsInvalidLatLng() bool {
	return utils.IsInvalidLatLng(q.Latitude, q.Longitude)
}
