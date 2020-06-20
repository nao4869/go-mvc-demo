package services

import (
	"fmt"
	"net/http"
	"sync"

	"../config"
	"../domain/github"
	"../domain/repositories"
	"../log"
	provider "../providers/github_provider"
	"../utils/errors"
)

// Entire Businness logic is in the services

type repositoryService struct{}

type repositoriesServiceInterface interface {
	CreateRepository(
		clientID string,
		request repositories.CreateRepositoryRequest,
	) (
		*repositories.CreateRepositoryResponse,
		errors.APIError,
	)

	CreateRepositories(
		request []repositories.CreateRepositoryRequest,
	) (
		repositories.CreateRepositoriesResponse,
		errors.APIError,
	)
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
func (s *repositoryService) CreateRepository(
	clientID string,
	input repositories.CreateRepositoryRequest,
) (
	*repositories.CreateRepositoryResponse,
	errors.APIError,
) {
	if error := input.Validate(); error != nil {
		return nil, error
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	log.Info(
		"about to send request to external API",
		fmt.Sprintf("client_id: %s", clientID),
		"status:pending",
	)

	// sending create repo request with secret access token
	response, error := provider.CreateRepository(
		config.GetGithubAccessToken(),
		request,
	)
	if error != nil {
		// new api error based on what we recieve from Github
		log.Info(
			"response obtained from external API",
			fmt.Sprintf("client_id: %s", clientID),
			"status:error",
		)

		return nil, errors.NewAPIError(
			error.StatusCode,
			error.Message,
		)
	}
	log.Info(
		"response obtained from external API",
		fmt.Sprintf("client_id: %s", clientID),
		"status:success",
	)

	result := repositories.CreateRepositoryResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

func (s *repositoryService) CreateRepositories(requests []repositories.CreateRepositoryRequest) (
	repositories.CreateRepositoriesResponse,
	errors.APIError,
) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateRepositoriesResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepositoryResults(&wg, input, output) // launch go routine

	// 3 requests to process
	// handle CreateRepository in a concurrent way
	for _, current := range requests {
		wg.Add(1)
		go s.CreateRepositoryConcurrent(current, input) // 3 go routines
	}

	// this will wait until go rotuine finish executing
	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0 // count for success response
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repositoryService) handleRepositoryResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateRepositoriesResponse) {
	var results repositories.CreateRepositoriesResponse

	for incomingEvent := range input {
		repositoryResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repositoryResult)

		// for go routine
		wg.Done()
	}
	output <- results
}

func (s *repositoryService) CreateRepositoryConcurrent(
	input repositories.CreateRepositoryRequest,
	output chan repositories.CreateRepositoriesResult,
) {
	// validating the request, if we have any error, send event to channel output
	if error := input.Validate(); error != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: error,
		}
		return
	}

	result, error := s.CreateRepository("TODO_client_id", input)
	if error != nil {
		output <- repositories.CreateRepositoriesResult{Error: error}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}
