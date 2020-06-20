package repositories

import (
	"log"
	"net/http"

	"../../domain/repositories"
	"../../services"
	"../../utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateRepository - create repo request
func CreateRepository(c *gin.Context) {
	// detect whether the incoming request is private or public
	// isPrivate := c.GetHeader("X-Private")

	var request repositories.CreateRepositoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientID := c.GetHeader("X-Client-Id")

	result, err := services.RepositoryService.CreateRepository(clientID, request)
	if err != nil {
		log.Println("BC")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// CreateRepositories -
func CreateRepositories(c *gin.Context) {
	var request []repositories.CreateRepositoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepositories(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
