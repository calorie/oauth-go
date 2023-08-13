package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/calorie/oauth-go/docs"
	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/usecase"
)

type TokenController struct {
	u *usecase.TokenUsecase
}

func NewTokenController(u *usecase.TokenUsecase) *TokenController {
	return &TokenController{
		u: u,
	}
}

// @Summary     post token
// @Description post token
// @Tags        token
// @Accept      json
// @Produce     json
// @Param       grant_type   query string true "Value MUST be set to authorization_code" Enums(authorization_code)
// @Param       code         query string true "The authorization code received from the authorization server"
// @Param       redirect_uri query string true "if the redirect_uri parameter was included in the authorization request as described in Section 4.1.1, and their values MUST be identical."
// @Param       client_id    query string true "if the client is not authenticating with the authorization server as described in Section 3.2.1"
// @Success     200 {object} domain.Token
// @Failure     400 {object} domain.HTTPError
// @Router      /token [post]
func (ac *TokenController) PostToken(c *gin.Context) {
	var r domain.TokenRequest

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPError{Error: domain.InvalidRequest, ErrorDescription: err.Error()})
		return
	}

	t, err := ac.u.PostToken(c, &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPError{Error: domain.InvalidRequest, ErrorDescription: err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}
