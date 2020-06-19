package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"../src/api/domain/repositories"
	"../src/api/services"
	"../src/api/utils/errors"
)

var (
	success = make(map[string]string, 0)
	failed  = make(map[string]errors.APIError, 0)
)

type createRepositoryResult struct {
	Request repositories.CreateRepositoryRequest
	Result  *repositories.CreateRepositoryResponse
	Error   errors.APIError
}

func getRequest() []repositories.CreateRepositoryRequest {
	result := make([]repositories.CreateRepositoryRequest, 0)

	file, error := os.Open("requests.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()

	// iterate file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		request := repositories.CreateRepositoryRequest{
			Name: line,
		}
		result = append(result, request)
	}

	return result
}

func main() {
	requests := getRequest()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepositoryResult)
	buffer := make(chan bool, 10) // create buffer
	var wg sync.WaitGroup

	go handleResults(&wg, input)

	// create each of requests in a linear way
	for _, request := range requests {
		buffer <- true // allow 10 eleements to send
		wg.Add(1)
		go createRepository(buffer, input, request)
	}
	wg.Wait() // wait until all channel input finished
	close(input)

	// write success or failed maps to disk or notify to the user via email or etc...
}

// iterating over the repository of results
// reason for passing pointer to WaitGroup is to ensure that always referencing same WaitGroup
func handleResults(wg *sync.WaitGroup, input chan createRepositoryResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done() // mark as done once finishing the process
	}
}

func createRepository(buffer chan bool, output chan createRepositoryResult, request repositories.CreateRepositoryRequest) {
	result, error := services.RepositoryService.CreateRepository(
		"your_client_id",
		request,
	)

	output <- createRepositoryResult{
		Request: request,
		Result:  result,
		Error:   error,
	}

	<-buffer
	//data := <-buffer // read from channel
}
