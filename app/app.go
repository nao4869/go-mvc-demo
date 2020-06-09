package app

import {
	"net/http"
	"github.com/nao4869/go-mvc-demo/controllers"
}

func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if error := http.ListenAndServe("localhost:8080", nil) {
		panic(error)
	}
}
