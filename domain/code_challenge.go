package domain

type CodeChallenge struct {
	AuthorizationCode   string                  `gorm:"primaryKey"`
	CodeChallenge       string                  `gorm:"not null"`
	CodeChallengeMethod CodeChallengeMethodType `gorm:"not null;type:char(4)"`
}
