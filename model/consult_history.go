package model

import "time"

// ConsultHistory ...
type ConsultHistory struct {
	ID        string `gorm:"primaryKey"`
	Status    string
	Action    string
	CreatedAt time.Time

	// Ref
	Question     ConsultQuestion `gorm:"foreignKey:QuestionID"`
	QuestionID   string
	Author       User `gorm:"foreignKey:AuthorID"`
	AuthorID     *string
	Specialist   User `gorm:"foreignKey:SpecialistID"`
	SpecialistID *string
	Staff        Staff `gorm:"foreignKey:StaffID"`
	StaffID      *string
}

// TableName overrides the table name
func (ConsultHistory) TableName() string {
	return "consult_histories"
}
