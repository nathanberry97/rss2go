package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/src/schema"
	"net/http"
)

func postRssFeed(router *gin.Engine) {
	router.POST("/rss_feed", func(ctx *gin.Context) {
		var body schema.RssPostBody
		ctx.BindJSON(&body)
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feed added successfully", "url": body.Url})
	})
}

func getRssFeeds(router *gin.Engine) {
	router.GET("/rss_feed", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feed fetched successfully"})
	})
}

func getRssFeedArticles(router *gin.Engine) {
	router.GET("/rss_feed/articles", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feed article fetched successfully"})
	})
}

func getRssFeedArticlesByFeed(router *gin.Engine) {
	router.GET("/rss_feed/:id/articles", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feed article fetched successfully"})
	})
}

func deleteRssFeed(router *gin.Engine) {
	router.DELETE("/rss_feed/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feed deleted successfully"})
	})
}
