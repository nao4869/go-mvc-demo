// package github

// import (
// 	"encoding/json"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateRequestAsJson(t *testing.T) {
// 	request := CreateRepoRequest{
// 		Name:        "test name",
// 		Description: "test repository for golang api demo",
// 		Homepage:    "https://github.com",
// 		Private:     true,
// 		HasIssues:   true,
// 		HasProjects: true,
// 		HasWiki:     true,
// 	}

// 	if request.Private {

// 	}

// 	// Marshal takes an input interface and attempts to create a valid json string.
// 	// 構造体のjsonタグがあれば、その値をキーとしてjson文字列を生成する
// 	// string(bytes) is the json github repo format data we retrieved
// 	bytes, err := json.Marshal(request)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, bytes)
// 	fmt.Println(string(bytes))
// 	assert.EqualValues(t, "test", string(bytes))

// 	var target CreateRepoRequest
// 	// Unmarshal takes an input byte array and a *pointer* that we're trying to fill using this json.
// 	err = json.Unmarshal(bytes, &target)
// 	assert.Nil(t, err)

// 	// それぞれのフィールドの値が同一かどうかの確認
// 	assert.EqualValues(t, target.Name, request.Name)
// 	assert.EqualValues(t, target.HasIssues, request.HasIssues)
// }
