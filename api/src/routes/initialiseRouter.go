package routes

import "github.com/gin-gonic/gin"

func InitialiseRouter() *gin.Engine {
	router := gin.Default()

	healthCheck(router)
	postRssFeed(router)
	getRssFeeds(router)
	getRssFeedArticles(router)
	getRssFeedArticlesByFeed(router)
	deleteRssFeed(router)

	return router
}
