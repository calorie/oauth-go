package domain

import (
	"time"
)

type AuthorizationCode struct {
	Code      string    `gorm:"primaryKey"`
	ExpiredAt time.Time `gorm:"not null"`
}
