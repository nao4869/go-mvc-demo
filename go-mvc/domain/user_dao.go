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

// GetUser - returns pointer to the user and an error
func GetUser(userID int64) (*User, *ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	// return the application error models when there is an error
	return nil, &ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
