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

func (r *AuthorizationCodeRepository) FindAuthorizationCode(code string) (*domain.AuthorizationCode, *gorm.DB) {
	ac := domain.AuthorizationCode{}
	db := r.db.Where("code = ? AND expired_at < ?", code, time.Now()).First(&ac)
	return &ac, db
}

func (r *AuthorizationCodeRepository) DeleteAuthorizationCode(code string) *gorm.DB {
	ac := domain.AuthorizationCode{
		Code: code,
	}
	return r.db.Delete(&ac)
}
