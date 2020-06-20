package app

import (
	"../controllers/repositories"
	"../log"
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
	log.Info(
		"about to start mapping the urls",
		"step:1",
		"status:pending",
	)

	router.POST("/repository", repositories.CreateRepository)
	router.POST("/repositories", repositories.CreateRepositories)

	log.Info(
		"urls successfully mapped",
		"step:2",
		"status:completed",
	)

	if error := router.Run(":8080"); error != nil {
		panic(error)
	}
}
