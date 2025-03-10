package routes

import "github.com/gin-gonic/gin"

func InitialiseRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("./web/templates/*.tmpl")

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
	getArticlesByFeedId(router)

	return router
}
