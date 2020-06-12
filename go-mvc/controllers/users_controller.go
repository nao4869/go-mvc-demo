package controllers

// requestから情報を受け取り、serviceへ送信する

import (
	"fmt"
	"net/http"
	"strconv"

	"../domain"
	"../services"
	"../utilities"
	"github.com/gin-gonic/gin"
)

// GetUser - allows to respond to same object in a json format and application/xml format
func GetUser(c *gin.Context) {
	// parama key is the url which is user_id in this case
	userID, error := (strconv.ParseInt(c.Param("user_id"), 10, 64))

	// Error hundling for user id format
	if error != nil {
		apiError := &domain.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		// responding in a JSON form to the client
		// c.JSON[http.StatusBadRequest, apiError]
		utilities.RespondError(c, apiError)
		return
	}
	fmt.Println(userID)

	// Error hundling to see whether user id exist or not
	user, apiError := services.GetUser(userID)
	if apiError != nil {
		utilities.RespondError(c, apiError)
		return
	}

	// return user to client
	utilities.Respond(c, http.StatusOK, user)
}
