package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nao4869/go-mvc-demo/src/api/controllers/repositories"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp -
func StartApp() {
	router.POST("/repositories", repositories.CreateRepository)

	if error := router.Run(":8080"); error != nil {
		panic(error)
	}
}
