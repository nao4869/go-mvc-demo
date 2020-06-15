package app

import "../controllers/repositories"

// mapUrls -
func mapUrls() {
	//router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepository)
}
