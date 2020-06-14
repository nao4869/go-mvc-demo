package app

import (
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
	mapUrls()

	if error := router.Run(":8080"); error != nil {
		panic(error)
	}
}
