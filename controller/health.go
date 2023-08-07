package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/calorie/oauth-go/docs"
)

type HealthController struct {}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// @Summary     get health
// @Description get health
// @Tags        health
// @Accept      json
// @Produce     json
// @Success     204 {object} domain.HTTPEmpty
// @Router      /v1/health [get]
func (ac *HealthController) GetHealth(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
