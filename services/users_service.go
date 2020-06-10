package services

import (
	"../domain"
)

// GetUser -
func GetUser(userID int64) (*domain.User, *domain.ApplicationError) {
	return domain.GetUser(userID)
}
