package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../../clients/restclient"
	"../../domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	createRepoRequestURL      = "https://api.github.com/user/repos"
)

// getAuthorizationHeader -
func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepository -
func CreateRepository(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}

	// headersにaceessTokenを指定する - 独自ヘッダー
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	// making post request
	response, error := restclient.Post(createRepoRequestURL, request, headers)

	// if we have an error, check error without checking success response
	if error != nil {
		log.Println(fmt.Sprintf("Error when creating new repository: %s", error.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    error.Error(),
		}
	}

	// reading the bytes of response body to check whether its valid or not
	bytes, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}
	// defer wait for function to return and close
	/* this will defer right before actual return in the function */
	defer response.Body.Close()

	// status code larger than 299 means error response
	if response.StatusCode > 299 {
		var errorResponse github.GithubErrorResponse
		if error := json.Unmarshal(bytes, &errorResponse); error != nil {
			// have error unmarsharing the response
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}

	// reaching here means that we have valid and succesfull response
	var result github.CreateRepoResponse
	if error := json.Unmarshal(bytes, &result); error != nil {
		log.Println(fmt.Sprintf("Error when trying to unmarshal for creating repository's successful response: %s", error.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when trying to unmarshal for creating repository's successful response",
		}
	}

	// return result for creating repository
	return &result, nil
}
