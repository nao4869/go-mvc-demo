package oauth

import (
	"net/http"

	"../../domain/oauth"
	"../../services"
	"../../utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateAccessToken -
func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest

	if error := c.ShouldBindJSON(&request); error != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(
			apiError.Status(),
			apiError,
		)
		return
	}

	token, error := services.OauthService.CreateAccessToken(request)
	if error != nil {
		c.JSON(error.Status(), error)
		return
	}
	c.JSON(http.StatusCreated, token)
}

// GetAccessToken -
func GetAccessToken(c *gin.Context) {
	tokenID := c.Param("token_id")
	token, error := services.OauthService.GetAccessToken(tokenID)

	if error != nil {
		c.JSON(error.Status(), error)
		return
	}
	c.JSON(http.StatusOK, token)
}
