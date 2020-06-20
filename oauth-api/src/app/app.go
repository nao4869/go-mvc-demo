package app

import (
	"../controllers/oauth"
	"../controllers/polo"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// StartApplication -
func StartApplication() {
	router = gin.Default()

	router.GET("/marco", polo.Polo)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)

	if error := router.Run(":8080"); error != nil {
		panic(error)
	}
}
