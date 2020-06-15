package config

import (
	"os"
)

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	// LogLevel -
	LogLevel      = "info"
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var (
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

// GetGithubAccessToken - function to return secret access token
func GetGithubAccessToken() string {
	return githubAccessToken
}

// IsProduction -
func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
