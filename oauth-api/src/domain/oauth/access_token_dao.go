package oauth

import (
	"fmt"

	"../../utils/errors"
)

var (
	tokens = make(
		map[string]*AccessToken,
		0,
	)
)

// Save -
func (at *AccessToken) Save() errors.APIError {
	at.AccessToken = fmt.Sprintf("&d", at.UserID)
	tokens[at.AccessToken] = at
	return nil
}

// GetAccessTokenByToken -
func GetAccessTokenByToken(accessToken string) (
	*AccessToken,
	errors.APIError,
) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundAPIError("no acceess token found with given parameters")
	}
	return token, nil
}
