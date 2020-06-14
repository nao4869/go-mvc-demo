package github_provider

import (
	"errors"
	//"fmt"
	"net/http"
	"testing"

	"../../../clients/restclient"
	"../../github"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("Invalid rest client response"),
	})

	response, error := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.EqualValues(t, "Invalid rest client response", error.Message)
}

// func TestDefer(t *testing.T) {
// 	// defer works as stack
// 	// 3 - 2 - 1
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	defer fmt.Println("3")

// 	fmt.Println("function's body")
// }
