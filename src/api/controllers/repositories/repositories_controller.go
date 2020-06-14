package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nao4869/go-mvc-demo/src/api/domain/repositories"
	"github.com/nao4869/go-mvc-demo/src/api/services"
	"github.com/nao4869/go-mvc-demo/src/api/utils/errors"
)

// CreateRepository - create repo request
func CreateRepository(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientID := c.GetHeader("X-Client-Id")

	result, err := services.RepositoryService.CreateRepo(clientID, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// CreateRepositories -
func CreateRepositories(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)
}
