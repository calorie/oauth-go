package main

import (
	"os"

	"github.com/calorie/oauth-go/controller"
	"github.com/calorie/oauth-go/repository"
	"github.com/calorie/oauth-go/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	db, err := repository.InitDb()
	if err != nil {
		panic("can't connect db")
	}

	cr := repository.NewClientRepositoty()
	sr := repository.NewScopeRepositoty()
	ac := controller.NewAuthorizeController(usecase.NewAuthorizeUsecase(cr, sr))
	hc := controller.NewHealthController()

	acr := repository.NewAuthorizationCodeRepositoty(db)
	ccr := repository.NewCodeChallengeRepositoty(db)
	ur := repository.NewUserRepositoty(db)
	dc := controller.NewDecisionController(usecase.NewDecisionUsecase(acr, ccr, ur))

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions("oauth", store))

	router.GET("/authorize", ac.GetAuthorize)
	router.POST("/decision", dc.PostDecision)

	v1 := router.Group("/v1")
	{

		v1.GET("/health", hc.GetHealth)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
