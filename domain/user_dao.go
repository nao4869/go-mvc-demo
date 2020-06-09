package domain

import (
	"errors"
	"fmt"
)

// temporaliry data base simulation
var (
	users = map[int64]*User{
		1: &User{ID: 1, FirstName: "test", LastName: "User", Email: "test@gmail.com"},
	}
)

// GetUser -
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("User %v was not found", userID))
}

// Another -
func Another() {
	user, integer, error := GetUser(1)

	if error != nil {
		return
	}
}
