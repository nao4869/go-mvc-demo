package domain

import (
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

	return nil, fmt.Errorf("User %v was not found", userID)
}
