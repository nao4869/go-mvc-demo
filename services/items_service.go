package services

import (
	"net/http"

	"../domain"
)

// GetItem -
func GetItem(itemID string) (*domain.Item, *domain.ApplicationError) {
	return nil, domain.ApplicationError{
		Message:    "to be implemented",
		StatusCode: http.StatusInternalServerError,
	}
}
