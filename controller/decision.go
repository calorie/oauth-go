package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/calorie/oauth-go/docs"
	"github.com/calorie/oauth-go/domain"
	"github.com/calorie/oauth-go/usecase"
)

type DecisionController struct {
	u *usecase.DecisionUsecase
}

func NewDecisionController(u *usecase.DecisionUsecase) *DecisionController {
	return &DecisionController{
		u: u,
	}
}

// @Summary     get decision
// @Description get decision
// @Tags        decision
// @Accept      x-www-form-urlencoded
// @Param       email    query string true "email" Format(email)
// @Param       password query string true "password"
// @Success     302
// @Header      302 {string} Location "https://client.example.org/cb?code=SplxlOBeZQQYbYS6WxSbIA&state=af0ifjsldkj"
// @Failure     400 {object} domain.HTTPError
// @Router      /decision [post]
func (ac *DecisionController) PostDecision(c *gin.Context) {
	var r domain.DecisionRequest

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPError{Error: domain.InvalidRequest, ErrorDescription: err.Error()})
		return
	}

	d, err := ac.u.PostDecision(c, &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPError{Error: domain.InvalidRequest, ErrorDescription: err.Error()})
		return
	}

	c.Redirect(http.StatusFound, d.Location)
}
