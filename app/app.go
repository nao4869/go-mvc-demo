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

// StartApplication -
func StartApplication() {
	mapUrls()

	if error := router.Run("localhost:8080", nil); error != nil {
		panic(error)
	}
}
