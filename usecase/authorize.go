package usecase

import (
	"errors"

	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/repository"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthorizeUsecase struct {
	cr *repository.ClientRepository
	sr *repository.ScopeRepository
}

func NewAuthorizeUsecase(cr *repository.ClientRepository, sr *repository.ScopeRepository) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		cr: cr,
		sr: sr,
	}
}

func (u *AuthorizeUsecase) GetAuthorize(c *gin.Context, r *domain.AuthorizeRequest) (*domain.Authorize, error) {
	client, err := u.cr.FindClient(r.ClientId)
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

	scopes := u.sr.FilterScope(r.Scope)

	session := sessions.Default(c)
	session.Set("client_id", client.ClientId)
	session.Set("state", r.State)
	session.Set("scope", u.sr.JoinScopes(scopes))
	session.Set("redirect_uri", uri)
	session.Set("code_challenge", r.CodeChallenge)
	session.Set("code_challenge_method", r.CodeChallengeMethod)
	session.Save()

	return &domain.Authorize{}, nil
}
