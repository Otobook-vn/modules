package model

import (
	"time"

	"gorm.io/datatypes"
)

// CarServiceCategory ...
type CarServiceCategory struct {
	ID           string `gorm:"primaryKey"`
	Name         datatypes.JSON
	Desc         datatypes.JSON
	SearchTokens TsVector
	Status       string
	Icon         datatypes.JSON
	CreatedAt    time.Time
}

// TableName overrides the table name
func (CarServiceCategory) TableName() string {
	return "car_service_categories"
}
