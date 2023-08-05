package main

import (
	"github.com/calorie/oauth-go/controller"
	_ "github.com/calorie/oauth-go/docs"
	"github.com/calorie/oauth-go/repository"
	"github.com/calorie/oauth-go/usecase"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @BasePath /v1

func main() {
	db, err := repository.InitDb()
	if err != nil {
		panic("can't connect db")
	}

	cr := repository.NewClientRepositoty()
	acr := repository.NewAuthorizationCodeRepositoty(db)
	ccr := repository.NewCodeChallengeRepositoty(db)
	ac := controller.NewAuthorizeController(usecase.NewAuthorizeUsecase(cr, acr, ccr))
	hc := controller.NewHealthController()

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/authorize", ac.GetAuthorize)

		v1.GET("/health", hc.GetHealth)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
