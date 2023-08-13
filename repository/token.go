package repository

import (
	"time"

	"github.com/calorie/oauth-go/domain"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepositoty(db *gorm.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) CreateToken(token string, userId string, clientId string, scope string, expiredIn time.Duration) *gorm.DB {
	ac := domain.AccessToken{
		Token: token,
		UserId: userId,
		ClientId: clientId,
		Scope: scope,
		ExpiredAt: time.Now().Add(expiredIn),
	}
	return r.db.Create(&ac)
}
