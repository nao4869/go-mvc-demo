package errors

import "net/http"

// APIError - interface
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"error,omitempty"`
}

// return - int
// takes pointer to *apiError
func (e *apiError) Status() int {
	return e.status
}

// return - int
// takes pointer to *apiError
func (e *apiError) Message() string {
	return e.message
}

// return - int
// takes pointer to *apiError
func (e *apiError) Error() string {
	return e.error
}

// NewAPIError -
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		status:  statusCode,
		message: message,
	}
}

// NewNotFoundAPIError -
func NewNotFoundAPIError(message string) APIError {
	return &apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}

// NewInternalServerError -
func NewInternalServerError(message string) APIError {
	return &apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

// NewBadRequestError -
func NewBadRequestError(message string) APIError {
	return &apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}
