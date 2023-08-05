package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func main() {
	router := gin.Default()
	router.GET("/health", health)

	router.Run(":8080")
}
