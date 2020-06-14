package app

import "github.com/nao4869/go-mvc-demo/src/api/controllers/repositories"

// mapUrls - 
func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepository)
}