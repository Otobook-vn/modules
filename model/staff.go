package model

import (
	"strings"
	"time"

	"github.com/Otobook-vn/modules/utils"
)

// Staff ...
type Staff struct {
	ID             string `gorm:"primaryKey"`
	Code           string
	Username       string
	Name           string
	Phone          string
	Email          string
	HashedPassword string
	Status         string
	SearchTokens   TsVector
	CreatedAt      time.Time

	// Ref
	Role   StaffRole `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE"`
	RoleID string
}

// TableName overrides the table name
func (Staff) TableName() string {
	return "staffs"
}

// GenerateSearchTokens ...
func (s *Staff) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(s.Name), s.Phone}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}

// IsRole ...
func (s Staff) IsRole(role string) bool {
	return s.RoleID == role
}

// HasStatus ...
func (s Staff) HasStatus(status string) bool {
	return s.Status == status
}
