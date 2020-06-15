package services

import (
	"strings"

	"../config"
	"../providers/github_provider"
	"../utils/errors"
	"../domain/repositories"
	"../domain/github"
)

// Entire Businness logic is in the services

type repositoryService struct{}

type repositoriesServiceInterface interface {
	CreateRepository(clientID string, request repositories.CreateRepositoryRequest) (*repositories.CreateRepositoryResponse, errors.APIError)
	//CreateRepositories(request []repositories.CreateRepositoryRequest) (repositories.CreateRepositoryResponse, errors.APIError)
}

var (
	// RepositoryService -
	RepositoryService repositoriesServiceInterface
)

func init() {
	RepositoryService = &repositoryService{}
}

/*
	return - CreateRepositoryResponse, errors.APIError
	argument - request interface{}
*/
func (s *repositoryService) CreateRepository(clientID string, input repositories.CreateRepositoryRequest) (*repositories.CreateRepositoryResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	// sending create repo request with secret access token
	response, error := provider.CreateRepository(config.GetGithubAccessToken(), request)
	if error != nil {
		// new api error based on what we recieve from Github
		return nil, errors.NewAPIError(error.StatusCode, error.Message)
	}

	result := repositories.CreateRepositoryResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}