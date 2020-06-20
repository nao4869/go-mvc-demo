package services

import (
	"time"

	"../domain/oauth"
	"../domain/user"
	"../utils/errors"
)

type oauthService struct{}

type oauthServiceInterface interface {
	/*
		argument - request oauth.AccessTokenRequest
		return type - oauth.AccessTokenRequest, errors.APIError
	*/
	CreateAccessToken(request oauth.AccessTokenRequest) (
		*oauth.AccessToken,
		errors.APIError,
	)
	GetAccessToken(accessToken string) (
		*oauth.AccessToken,
		errors.APIError,
	)
}

var (
	// OauthService -
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (
	*oauth.AccessToken,
	errors.APIError,
) {
	if error := request.Validate(); error != nil {
		return nil, error
	}

	user, error := user.GetUserByUsernameAndPassword(
		request.Username,
		request.Password,
	)
	if error != nil {
		return nil, error
	}

	token := oauth.AccessToken{
		UserID:  user.ID,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}

	// saving access token
	if error := token.Save(); error != nil {
		return nil, error
	}
	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (
	*oauth.AccessToken,
	errors.APIError,
) {
	token, error := oauth.GetAccessTokenByToken(accessToken)
	if error != nil {
		return nil, error
	}

	if token.IsExpired() {
		return nil, errors.NewNotFoundAPIError("no access token found for given information")
	}
	return token, error
}
