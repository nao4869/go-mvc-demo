package services

import (
	"net/http"

	"../domain"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemID string) (*domain.Item, *domain.ApplicationError) {
	return nil, &domain.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
