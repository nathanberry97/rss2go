package routes

import "github.com/gin-gonic/gin"

func InitialiseRouter(cssFile string) *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("./web/templates/*.tmpl")

	// Health check
	healthCheck(router)

	// HTML routes
	articlesByFeedPage(router, cssFile)
	articlesFavourite(router, cssFile)
	articlesPage(router, cssFile)
	articlesReadLater(router, cssFile)
	feedsPage(router, cssFile)
	settings(router, cssFile)

	// Article routes
	getArticles(router)
	getArticlesByFeedId(router)

	// Feed routes
	deleteFeed(router)
	getFeeds(router)
	postFeed(router)
	postFeedOpml(router)

	// Read later routes
	deleteReadLater(router)
	getReadLater(router)
	postReadLater(router)

	// Favourite routes
	deleteFavourite(router)
	getFavourites(router)
	postFavourite(router)

	return router
}
