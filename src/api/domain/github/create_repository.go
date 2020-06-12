package github

// CreateRepoRequest - fields are specified at Github API
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse -
type CreateRepoResponse struct {
	ID          int64          `json:"id"`
	Name        string         `json:"string"`
	FullName    string         `json:"string"`
	Owner       RepoOwner      `json:"owner"`
	Permissions RepoPermissons `json:"permissions"`
}

// RepoIwner -
type RepoIwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// RepoPermissions -
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	Haspull bool `json:"push"`
	HasPush bool `json:"pull"`
}
