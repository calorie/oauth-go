package domain

import (
	"time"
)

type AuthorizationCode struct {
	Code        string    `gorm:"primaryKey"`
	UserId      string    `gorm:"not null"`
	ClientId    string    `gorm:"not null"`
	Scope       string    `gorm:"not null"`
	RedirectUri string    `gorm:"not null"`
	ExpiredAt   time.Time `gorm:"not null"`
}
