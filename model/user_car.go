package model

import "time"

// UserCar ...
type UserCar struct {
	ID           string `gorm:"primaryKey"`
	Year         string
	LicensePlate string
	VINNumber    string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Ref
	User      User `gorm:"UserID"`
	UserID    string
	Brand     CarBrand `gorm:"foreignKey:BrandID"`
	BrandID   *string
	Model     CarModel `gorm:"foreignKey:ModelID"`
	ModelID   *string
	Version   CarModelVersion `gorm:"foreignKey:VersionID"`
	VersionID *string
	Color     CarColor `gorm:"foreignKey:ColorID"`
	ColorID   *string
}

// TableName overrides the table name
func (UserCar) TableName() string {
	return "user_cars"
}
