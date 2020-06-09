package controllers

// requestから情報を受け取り、serviceへ送信する

import (
	"net/http"
	"strconv"
	"github.com/nao4869/go-mvc-demo/services"
)

// GetUser -
func GetUser(response http.ResponseWriter, request *http.Request) {
	userID, error := (strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64))

	if error != nil {
		return
	}

	user, error := services.GetUser(userID)
	if error != nil {
		// handle the error and return to the client
		return
	}

}
