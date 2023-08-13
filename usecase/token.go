package usecase

import (
	"errors"
	"time"

	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TokenUsecase struct {
	acr *repository.AuthorizationCodeRepository
	tr *repository.TokenRepository
	ccr *repository.CodeChallengeRepository
}

func NewTokenUsecase(acr *repository.AuthorizationCodeRepository, tr *repository.TokenRepository, ccr *repository.CodeChallengeRepository) *TokenUsecase {
	return &TokenUsecase{
		acr: acr,
		tr: tr,
		ccr: ccr,
	}
}

func (u *TokenUsecase) PostToken(c *gin.Context, r *domain.TokenRequest) (*domain.Token, error) {
	code, db := u.acr.FindAuthorizationCode(r.Code)
	if db.Error != nil {
		return nil, db.Error
	}

	if r.RedirectUri != code.RedirectUri {
		return nil, errors.New("redirect_uri is wrong")
	}

	if r.ClientId != code.ClientId {
		return nil, errors.New("client_id is wrong")
	}

	token := uuid.NewString()
	expiredIn := 24 * time.Hour
	db = u.tr.CreateToken(token, code.UserId, code.ClientId, code.Scope, expiredIn)
	if db.Error != nil {
		return nil, db.Error
	}

	cc, db := u.ccr.FindCodeChallenge(code.Code)
	if db.Error != nil {
		return nil, db.Error
	}

	if cc.Verify(r.CodeVerifier) {
		return nil, errors.New("code_verifier is wrong")
	}

	// db = u.acr.DeleteAuthorizationCode(r.Code)
	// if db.Error != nil {
	// 	return nil, db.Error
	// }

	t := &domain.Token{
		AccessToken: token,
		TokenType: "Bearer",
		ExpiredIn: int(expiredIn / time.Second),
	}

	return t, nil
}
