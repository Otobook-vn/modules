package model

import (
	"time"

	"gorm.io/datatypes"
)

// ConsultComment ...
type ConsultComment struct {
	ID        string `gorm:"primaryKey"`
	Content   string
	Photos    datatypes.JSON
	CreatedAt time.Time

	// Ref
	Author     User `gorm:"foreignKey:AuthorID"`
	AuthorID   string
	Question   ConsultQuestion `gorm:"foreignKey:QuestionID"`
	QuestionID string
}

// TableName overrides the table name
func (ConsultComment) TableName() string {
	return "consult_comments"
}
