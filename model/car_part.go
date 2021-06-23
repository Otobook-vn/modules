package model

import (
	"time"

	"gorm.io/datatypes"
)

// CarPart ...
type CarPart struct {
	ID           string `gorm:"primaryKey"`
	Name         datatypes.JSON
	SearchTokens TsVector
	Status       string
	IsOther      bool
	CreatedAt    time.Time
}

// TableName overrides the table name
func (CarPart) TableName() string {
	return "car_parts"
}
