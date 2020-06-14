package app

import "github.com/nao4869/go-mvc-demo/src/api/controllers/repositories"

// mapUrls - 
func mapUrls() {
	router.GET("/marco", polo.Polo)

	//router.POST("/repository", repositories.CreateRepo)

	// any post request maches /repositories will be handled by repositories_controller.go
	router.POST("/repositories", repositories.CreateRepository)
}