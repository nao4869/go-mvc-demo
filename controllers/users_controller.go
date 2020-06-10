package controllers

// requestから情報を受け取り、serviceへ送信する

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetUser -
func GetUser(response http.ResponseWriter, request *http.Request) {
	userID, error := (strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64))

	if error != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("user id must be a number"))
		return
	}

	user, error := services.GetUser(userID)
	if error != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(error.Error()))
		// handle the error and return to the client
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)

}
