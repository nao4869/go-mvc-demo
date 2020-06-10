package services

import (
	"github.com/nao4869/go-mvc-demo/domain"
)

// GetUser -
func GetUser(userID int64) (*domain.User, *domain.ApplicationError) {
	return domain.GetUser(userID)
}
