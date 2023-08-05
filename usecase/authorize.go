package usecase

import (
	"errors"
	"net/url"

	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorizeUsecase struct {
	cr  *repository.ClientRepository
	acr *repository.AuthorizationCodeRepository
	ccr *repository.CodeChallengeRepository
}

func NewAuthorizeUsecase(cr *repository.ClientRepository, acr *repository.AuthorizationCodeRepository, ccr *repository.CodeChallengeRepository) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		cr: cr,
		acr: acr,
		ccr: ccr,
	}
}

func (u *AuthorizeUsecase) GetAuthorize(c *gin.Context, r *domain.AuthorizeRequest) (*domain.Authorize, error) {
	client, err := domain.FindClient(r.ClientId)
	if err != nil {
		return nil, err
	}

	uri, err := client.FindRedirectUri(r.RedirectUri)
	if err != nil {
		return nil, err
	}

	if !r.ResponseTypeCode() {
		return nil, errors.New("response_type is wrong")
	}

	if !r.CodeChallengeMethodS256() {
		return nil, errors.New("code_challenge_method is wrong")
	}

	domain.FilterScope(r.Scope)

	code := uuid.NewString()
	db := u.acr.CreateAuthorizationCode(code)
	if db.Error != nil {
		return nil, db.Error
	}

	db = u.ccr.CreateCodeChallenge(code, r.CodeChallenge, r.CodeChallengeMethod)
	if db.Error != nil {
		return nil, db.Error
	}

	a := &domain.Authorize{Location: u.location(uri, code, r.State)}
	return a, nil
}

func (u *AuthorizeUsecase) location(uri string, code string, state string) string {
	url, err := url.Parse(uri)
	if err != nil {
		panic("can't parse uri")
	}

	q := url.Query()
	q.Set("code", code)
	q.Set("state", state)
	url.RawQuery = q.Encode()

	return url.String()
}
