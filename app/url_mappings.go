package app

import (
	"../controllers"
)

func mapUrls() {
	router.GET("/users/:user_id?caller=1234", controllers.GetUser)
}
