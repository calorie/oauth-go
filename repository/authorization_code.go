package repository

import (
	"time"

	"github.com/calorie/oauth-go/domain"
	"gorm.io/gorm"
)

type AuthorizationCodeRepository struct {
	db *gorm.DB
}

func NewAuthorizationCodeRepositoty(db *gorm.DB) *AuthorizationCodeRepository {
	return &AuthorizationCodeRepository{
		db: db,
	}
}

func (r *AuthorizationCodeRepository) CreateAuthorizationCode(code string) *gorm.DB {
	expiredAt := time.Now().Add(10 * time.Minute)
	return r.db.Create(&domain.AuthorizationCode{Code: code, ExpiredAt: expiredAt})
}
