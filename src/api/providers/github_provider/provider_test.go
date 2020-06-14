package provider

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"net/http"
	"testing"

	"../../clients/restclient"
	"../github"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})

	response, err := CreateRepository("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid restclient response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	//restclient.StartMockups()
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("Invalid rest client response"),
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})

	response, error := CreateRepository("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusInternalServerError, error.Message)
	assert.EqualValues(t, "Invalid rest client response", error.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, error := CreateRepository("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusInternalServerError, error.Message)
	assert.EqualValues(t, "Invalid json response body", error.Message)
}

func TestCreateUnauthorizedError(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authentication"}`)),
		},
	})

	response, error := CreateRepository("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusUnauthorized, error.Message)
	assert.EqualValues(t, "Invalid json response body", error.Message)
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "1"}`)),
		},
	})

	response, error := CreateRepository("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusInternalServerError, error.Message)
	assert.EqualValues(t, "Error when trying to unmarshal for creating repository's successful response", error.Message)
}

// created repo with no error
func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "1"}`)),
		},
	})

	response, error := CreateRepository("", github.CreateRepoRequest{})
	assert.Nil(t, error)
	assert.NotNil(t, response)
	assert.EqualValues(t, 1, response.ID)
	assert.EqualValues(t, "golang-tutorial", response.Name)
	assert.EqualValues(t, "golang-tutorial", response.FullName)
}

// func TestDefer(t *testing.T) {
// 	// defer works as stack
// 	// 3 - 2 - 1
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	defer fmt.Println("3")

// 	fmt.Println("function's body")
// }
