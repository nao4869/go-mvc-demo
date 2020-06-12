package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
only the function will be tested. In this case, GetUser function will be testes
when running go test command, go test -cover etc.
*/

func TestGetUserNoUserFound(t *testing.T) {
	// Execution
	user, error := GetUser(0)

	// Validation
	assert.Nil(t, user, "Fail: user id is 0")
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusNotFound, error.StatusCode)
	assert.EqualValues(t, "not found", error.Code)
	assert.EqualValues(t, "user 0 does not exist", error.Message)

	/* these error hundling using if statement can be replaced with assert */
	// if user != nil {
	// 	t.Error("Fail: user id is 0")
	// }

	// if error == nil {
	// 	t.Error("Expecting: when user id is 0, the error should occur")
	// }

	// if error.StatusCode != http.StatusNotFound {
	// 	t.Error("Expecting: expecting http code of 404 when user is not found")
	// }
}

/* TestGetUserNoError - testing to see
user id is not nil
user id, first name, last name, email is equal to the value we have created
*/
func TestGetUserNoError(t *testing.T) {
	user, error := GetUser(1)

	assert.Nil(t, error)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Test", user.FirstName)
	assert.EqualValues(t, "User", user.LastName)
	assert.EqualValues(t, "test@gmail.com", user.Email)
}
