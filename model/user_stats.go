package model

// UserStats ...
type UserStats struct {
	ID               string `gorm:"primaryKey"`
	Post             int
	CurrentCoin      int
	Reward           int
	Car              int
	Follower         int
	Following        int
	ReferralTotal    int
	ReferralSuccess  int
	ConsultantRating float64

	// Ref
	User   User `gorm:"foreignKey:UserID"`
	UserID string
}

// TableName overrides the table name
func (UserStats) TableName() string {
	return "user_stats"
}
