package repositories

// CreateRepositoryRequest -
type CreateRepositoryRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

// CreateRepositoryResponse -
type CreateRepositoryResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}
