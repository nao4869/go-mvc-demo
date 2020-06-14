package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

// Mock -
type Mock struct {
	URL        string
	HTTPMethod string
	Response  *http.Response
	Err        error
}

// getMockID -
func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// StartMockups -
func StartMockups() {
	enabledMocks = true
}

// StopMockups -
func StopMockups() {
	enabledMocks = false
}

// AddMockup -
func AddMockup(mock Mock) {
	mocks[getMockID(mock.HTTPMethod,mock.URL)] = &mock
}

// FlushMockups -
func FlushMockups() {
	mocks = make(map[string]*Mock)
}

// Post -
// function to create Post request to the github api
// if we have valid jsonBytes -> create request -> create client
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockID(http.MethodPost, url)]

		// if mocks for given request does not exist, return error.
		if mock == nil {
			return nil, errors.New("No mockup found for this request")
		}
		return mock.Response, mock.Err
	}

	// responseのbodyにエラーがある場合は、できることはない為、errorをreturn
	jsonBytes, error := json.Marshal(body)
	if error != nil {
		return nil, error
	}
	request, error := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
