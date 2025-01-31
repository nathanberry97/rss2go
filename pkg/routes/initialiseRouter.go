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

	// RSS feed routes
	postRssFeed(router)
	getRssFeeds(router)
	deleteRssFeed(router)

	// RSS article routes
	GetRssArticles(router)

	return router
}
