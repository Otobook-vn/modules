package model

import "time"

// ConsultQuestionLike ...
type ConsultQuestionLike struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time

	// Ref
	User       User `gorm:"foreignKey:UserID"`
	UserID     string
	Question   ConsultQuestion `gorm:"foreignKey:QuestionID"`
	QuestionID string
}

// TableName overrides the table name
func (ConsultQuestionLike) TableName() string {
	return "consult_question_likes"
}
