package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/calorie/oauth-go/docs"
	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/usecase"
)

type AuthorizeController struct {
	u *usecase.AuthorizeUsecase
}

func NewAuthorizeController(u *usecase.AuthorizeUsecase) *AuthorizeController {
	return &AuthorizeController{
		u: u,
	}
}

// @Summary     get authorize
// @Description get authorize
// @Tags        authorize
// @Param       response_type         query string true "code for Authorization Code Grant" Enums(code)
// @Param       client_id             query string true "OAuth 2.0 Client Identifier valid at the Authorization Server."
// @Param       redirect_uri          query string true "Redirection URI to which the response will be sent."
// @Param       scope                 query string true "OpenID Connect requests MUST contain the openid scope value."
// @Param       state                 query string true "Opaque value used to maintain state between the request and the callback."
// @Param       code_challenge        query string true "https://datatracker.ietf.org/doc/html/rfc7636"
// @Param       code_challenge_method query string true "https://datatracker.ietf.org/doc/html/rfc7636" Enums(S256)
// @Success     200
// @Router      /authorize [get]
func (ac *AuthorizeController) GetAuthorize(c *gin.Context) {
	var r domain.AuthorizeRequest

	err := c.ShouldBindQuery(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPError{Error: domain.InvalidRequest, ErrorDescription: err.Error()})
		return
	}

	_, err = ac.u.GetAuthorize(c, &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPError{Error: domain.InvalidRequest, ErrorDescription: err.Error()})
		return
	}

	c.HTML(http.StatusOK, "authorize/index.tmpl", gin.H{})
}
