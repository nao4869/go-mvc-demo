package utilities

import (
	"../domain"
	"github.com/gin-gonic/gin"
)

// Respond -
func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

// RespondError -
func RespondError(c *gin.Context, err *domain.ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)
}
