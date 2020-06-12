package github

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestCreateRequestAsJson(t *testing.T) {
	request := CreateRepoRequest {
		Name: "test repository for golang api demo",
		Description: "test repository for golang api demo",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: true,
		HasProjects: true,
		HasWiki: true,
	}

	// Marshal takes an input interface and attempts to create a valid json string.
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	// Unmarshal takes an input byte array and a *pointer* that we're trying to fill using this json.
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}