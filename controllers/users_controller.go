package controllers

// requestから情報を受け取り、serviceへ送信する

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../domain"
	"../services"
	"../utilities"
	"github.com/gin-gonic/gin"
)

// GetUser -
func GetUser(c *gin.Context) {
	// parama key is the url which is user_id in this case
	userID, error := (strconv.ParseInt(c.Param("user_id"), 10, 64))

	// Error hundling for user id format
	if error != nil {
		apiError := &domain.ApplicationError{
			Message:    "user id was must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		// responding in a JSON form to the client
		// c.JSON[http.StatusBadRequest, apiError]
		utilities.Respond(c, apiError.StatusCode, apiError)
		return
	}

	// Error hundling to see whether user id exist or not
	user, error := services.UserService.GetUser(userID)
	if error != nil {
		c.JSON[apiError.StatusCode, apiError]
		return
	}

	// return user to client
	//c.JSON(http.StatusOK, user)
	utilities.Respond(c, http.StatusOK, user)
}