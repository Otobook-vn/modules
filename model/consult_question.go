package model

import (
	"strings"
	"time"

	"gorm.io/datatypes"

	"github.com/Otobook-vn/modules/utils"
)

// ConsultQuestion ...
type ConsultQuestion struct {
	ID                      string `gorm:"primaryKey"`
	Title                   string
	SearchTokens            TsVector
	Content                 string
	Status                  string
	Photos                  datatypes.JSON
	HasPublicStatus         bool // Mark this question status is one of these statuses: open, completed, resolved
	IsHasSpecialistInCharge bool // Mark this question has specialist in charged or not
	Rating                  int
	StatsView               int
	StatsLike               int
	StatsComment            int
	ClosedReason            string
	RejectReason            string
	CreatedAt               time.Time
	UpdatedAt               time.Time
	RatedAt                 time.Time

	// Ref
	Author       User `gorm:"foreignKey:AuthorID"`
	AuthorID     string
	Specialist   User `gorm:"foreignKey:SpecialistID"`
	SpecialistID *string
	Category     ConsultCategory `gorm:"foreignKey:CategoryID"`
	CategoryID   string
}

// TableName overrides the table name
func (ConsultQuestion) TableName() string {
	return "consult_questions"
}

// GenerateSearchTokens ...
func (s *ConsultQuestion) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(s.Title)}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
