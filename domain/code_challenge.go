package domain

type CodeChallenge struct {
	AuthorizationCode   string                  `gorm:"primaryKey"`
	CodeChallenge       string                  `gorm:"not null"`
	CodeChallengeMethod CodeChallengeMethodType `gorm:"not null;type:char(4)"`
}

func (c *CodeChallenge) Verify(verifier string) bool {
	return true
}
