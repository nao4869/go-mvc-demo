package github_provider

import (
	"fmt"

	"../../github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, "abc123")
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (github.CreateRepoResponse, github.GithubErrorResponse) {
	header := getAuthorizationHeader(accessToken)
	fmt.Println(header)

	return github.CreateRepoResponse{}, github.GithubErrorResponse{}
}
