package model

import (
	"time"

	"gorm.io/datatypes"
)

// CarTransmission ...
type CarTransmission struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Status    string
	CreatedAt time.Time
}

// TableName overrides the table name
func (CarTransmission) TableName() string {
	return "car_transmissions"
}

// CarBodyStyle ...
type CarBodyStyle struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Status    string
	CreatedAt time.Time
}

// TableName overrides the table name
func (CarBodyStyle) TableName() string {
	return "car_body_styles"
}

// CarColor ...
type CarColor struct {
	ID        string
	Color     string `gorm:"primaryKey"`
	Name      datatypes.JSON
	CreatedAt time.Time
}

// TableName overrides the table name
func (CarColor) TableName() string {
	return "car_colors"
}
