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

func (r *AuthorizationCodeRepository) CreateAuthorizationCode(code string, userId string, clientId string, scope string, redirectUri string) *gorm.DB {
	ac := domain.AuthorizationCode{
		Code: code,
		UserId: userId,
		ClientId: clientId,
		Scope: scope,
		RedirectUri: redirectUri,
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}
	return r.db.Create(&ac)
}
