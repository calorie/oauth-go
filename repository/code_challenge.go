package repository

import (
	"github.com/calorie/oauth-go/domain"
	"gorm.io/gorm"
)

type CodeChallengeRepository struct {
	db *gorm.DB
}

func NewCodeChallengeRepositoty(db *gorm.DB) *CodeChallengeRepository {
	return &CodeChallengeRepository{
		db: db,
	}
}

func (r *CodeChallengeRepository) CreateCodeChallenge(code string, cc string, ccm domain.CodeChallengeMethodType) *gorm.DB {
	return r.db.Create(&domain.CodeChallenge{AuthorizationCode: code, CodeChallenge: cc, CodeChallengeMethod: ccm})
}

func (r *CodeChallengeRepository) FindCodeChallenge(code string) (*domain.CodeChallenge, *gorm.DB) {
	cc := domain.CodeChallenge{}
	db := r.db.Where("authorization_code = ?", code).First(&cc)
	return &cc, db
}
