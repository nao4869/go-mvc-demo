package errors

import "net/http"

// APIError - interface
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AError   string `json:"error,omitempty"`
}

// return - int
// takes pointer to *apiError
func (e *apiError) Status() int {
	return e.AStatus
}

// return - int
// takes pointer to *apiError
func (e *apiError) Message() string {
	return e.AMessage
}

// return - int
// takes pointer to *apiError
func (e *apiError) Error() string {
	return e.AError
}

// NewAPIError -
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}

// NewNotFoundAPIError -
func NewNotFoundAPIError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

// NewInternalServerError -
func NewInternalServerError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

// NewBadRequestError -
func NewBadRequestError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
