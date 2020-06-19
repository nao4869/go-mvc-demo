package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"../clients/restclient"
	"../domain/repositories"
	"github.com/nao4869/go-mvc-demo/src/api/services"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

// invalid input name for CreateRepository function
func TestCreateRepositoryInvalidInputName(t *testing.T) {
	request := repositories.CreateRepositoryRequest{}

	result, error := services.RepositoryService.CreateRepository(
		"Client ID",
		request,
	)

	assert.Nil(t, result)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusBadRequest, error.Status())
	assert.EqualValues(t, "Invalid json body", error.Message())
}

// for invalid response from github
func TestCreateRepositoryErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})
	request := repositories.CreateRepositoryRequest{Name: "testing"}
	result, error := RepositoryService.CreateRepository(request)

	assert.Nil(t, result)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusBadRequest, error.Status())
	assert.EqualValues(t, "Requires authentication", error.Message())
}

func TestCreateRepositoryNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1, "name": "testing", "owner": {"login": "nao4869"}}`)),
		},
	})
	request := repositories.CreateRepositoryRequest{Name: "testing"}
	result, error := RepositoryService.CreateRepository(request)

	assert.Nil(t, result)
	assert.NotNil(t, error)
	assert.EqualValues(t, 1, result.ID)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepositoryRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repositoryService{}

	go service.CreateRepositoryConcurrent(
		request,
		output,
	)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}
