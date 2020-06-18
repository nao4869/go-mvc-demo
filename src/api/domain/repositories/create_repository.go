package repositories

import (
	"strings"

	"../../utils/errors"
)

// CreateRepositoryRequest -
type CreateRepositoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate -
func (r *CreateRepositoryRequest) Validate()errors.APIError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

// CreateRepositoryResponse -
type CreateRepositoryResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

// CreateRepositoriesResponse -
type CreateRepositoriesResponse struct {
	StatusCode int                        `json:"status"`
	Results    []CreateRepositoriesResult `json:"results"`
}

// CreateRepositoriesResult -
type CreateRepositoriesResult struct {
	Response *CreateRepositoryResponse `json:"repo"`
	Error    errors.APIError          `json:"error"`
}
