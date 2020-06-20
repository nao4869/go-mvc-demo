package oauth

import (
	"strings"

	"github.com/nao4869/go-mvc-demo/src/api/utils/errors"
)

// AccessTokenRequest -
type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate -
func (r *AccessTokenRequest) Validate() errors.APIError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid user name")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
