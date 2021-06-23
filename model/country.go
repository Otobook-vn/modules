package model

import (
	"time"

	"gorm.io/datatypes"
)

// Country ...
type Country struct {
	ID        string
	Code      string `gorm:"primaryKey"`
	Name      datatypes.JSON
	CreatedAt time.Time
}

// TableName overrides the table name
func (Country) TableName() string {
	return "countries"
}
