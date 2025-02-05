package routes

import "github.com/gin-gonic/gin"

func InitialiseRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	// Health check
	healthCheck(router)

	// HTML routes
	articlesPage(router)
	feedsPage(router)

	// Feed routes
	getFeeds(router)
	postFeed(router)
	deleteFeed(router)

	// Article routes
	getArticles(router)

	return router
}
