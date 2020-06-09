package controllers

import (
	"log"
	"net/http"
)

// GetUser -
func GetUser(response http.ResponseWriter, request *http.Request) {
	userID := request.URL.Query().Get("user_id")
	log.Printf("user_id %v", userID)
}
