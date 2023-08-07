package usecase

import (
	"net/url"

	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/repository"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DecisionUsecase struct {
	acr *repository.AuthorizationCodeRepository
	ccr *repository.CodeChallengeRepository
	ur *repository.UserRepository
}

func NewDecisionUsecase(acr *repository.AuthorizationCodeRepository, ccr *repository.CodeChallengeRepository, ur *repository.UserRepository) *DecisionUsecase {
	return &DecisionUsecase{
		acr: acr,
		ccr: ccr,
		ur: ur,
	}
}

func (u *DecisionUsecase) PostDecision(c *gin.Context, r *domain.DecisionRequest) (*domain.Decision, error) {
	session := sessions.Default(c)
	clientId := session.Get("client_id").(string)
	redirectUri := session.Get("redirect_uri").(string)
	scope := session.Get("scope").(string)
	state := session.Get("state").(string)
	codeChallenge := session.Get("code_challenge").(string)
	codeChallengeMethod := session.Get("code_challenge_method").(domain.CodeChallengeMethodType)
	session.Clear()

	user, db := u.ur.FindUser(r.Email, r.Password)
	if db.Error != nil {
		return nil, db.Error
	}

	code := uuid.NewString()
	db = u.acr.CreateAuthorizationCode(code, user.Id, clientId, scope, redirectUri)
	if db.Error != nil {
		return nil, db.Error
	}

	db = u.ccr.CreateCodeChallenge(code, codeChallenge, codeChallengeMethod)
	if db.Error != nil {
		return nil, db.Error
	}

	d := &domain.Decision{Location: u.location(redirectUri, code, state)}
	return d, nil
}

func (u *DecisionUsecase) location(uri string, code string, state string) string {
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
