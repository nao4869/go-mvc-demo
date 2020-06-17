package app

import "../controllers/repositories"

// mapUrls -
func mapUrls() {
	//router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepository)
	router.POST("/repositories", repositories.CreateRepositories)
}
