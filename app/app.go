package app

import (
	"net/http"

	"github.com/nao4869/go-mvc-demo/controllers"
)

// StartApplication -
func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if error := http.ListenAndServe("localhost:8080", nil); error != nil {
		panic(error)
	}
}
