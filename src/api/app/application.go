package app

import (
	"../controllers/repositories"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp -
func StartApp() {
	router.POST("/repository", repositories.CreateRepository)
	router.POST("/repositories", repositories.CreateRepositories)

	if error := router.Run(":8080"); error != nil {
		panic(error)
	}
}
