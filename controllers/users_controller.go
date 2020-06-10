package controllers

// requestから情報を受け取り、serviceへ送信する

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nao4869/go-mvc-demo/domain"
)

// GetUser -
func GetUser(response http.ResponseWriter, request *http.Request) {
	userID, error := (strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64))

	// Error hundling for user id format
	if error != nil {
		arrError := &domain.ApplicationError{
			Message:    "user id was must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(arrError)
		response.WriteHeader(arrError.StatusCode)
		response.Write([]byte("user id must be a number"))
		return
	}

	// Error hundling to see whether user id exist or not
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
