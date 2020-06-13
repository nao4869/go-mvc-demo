package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/*
	function to create Post request to the github api
	 - if we have valid jsonBytes -> create request -> create client
*/
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
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
