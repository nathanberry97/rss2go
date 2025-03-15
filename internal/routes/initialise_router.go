package routes

import "github.com/gin-gonic/gin"

func InitialiseRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("./web/templates/*.tmpl")

	// Health check
	healthCheck(router)

	// HTML routes
	articlesByFeedPage(router)
	articlesFavorite(router)
	articlesPage(router)
	articlesReadLater(router)
	feedsPage(router)
	settings(router)

	// Article routes
	getArticles(router)
	getArticlesByFeedId(router)

	// Feed routes
	deleteFeed(router)
	getFeeds(router)
	postFeed(router)

	// Read later routes
	deleteReadLater(router)
	getReadLater(router)
	postReadLater(router)

	return router
}
