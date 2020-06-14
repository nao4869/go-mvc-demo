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
	// type CreateRepositoryRequest empty variable
	var request repositories.CreateRepositoryRequest

	// if the JSON value in request is valie - populate avobe createRepoRequest
	if error := c.ShouldBindJSON(&request); error != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}

	result, error := services.RepositoryService.CreateRepository(request)
	if error != nil {
		c.JSON(apiError.Status(), apiError)
		return
	}

	c.JSON(http.StatusCreated, result)
}
