package user

import (
	"fmt"

	"../../utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"test": &User{
			ID:       1,
			Username: "testUser",
		},
	}
)

// GetUserByUsernameAndPassword -
func GetUserByUsernameAndPassword(
	username string,
	password string,
) (*User, errors.APIError) {

	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundAPIError(
			fmt.Sprintf(
				"user with username %s not found",
				username,
			),
		)
	}
	return user, nil
}
