package domain

import (
	"fmt"
	"net/http"
)

// temporaliry data base simulation
var (
	users = map[int64]*User{
		1: &User{ID: 1, FirstName: "test", LastName: "User", Email: "test@gmail.com"},
	}
)

// GetUser -
func GetUser(userID int64) (*User, *ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	// return the application error models when there is an error
	return nil, &ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
