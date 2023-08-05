package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/calorie/oauth-go/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @BasePath /v1

// Health godoc
// @Summary     get health
// @Description get health
// @Tags        health
// @Accept      json
// @Produce     json
// @Success     204 {object} domain.HTTPEmpty
// @Router      /health [get]
func health(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/health", health)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
