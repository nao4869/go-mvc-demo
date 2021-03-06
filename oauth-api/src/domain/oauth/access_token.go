package oauth

import (
	"time"
)

// AccessToken -
type AccessToken struct {
	AccessToken string    `json:"access_token"`
	UserID      int64     `json:"user_id"`
	Expires     int64 `json:"expires"`
}

// IsExpired -
func (at *AccessToken) IsExpired() bool {
	return time.Unix(
		at.Expires, 
		0,
	).UTC().Before(time.Now().UTC())
}