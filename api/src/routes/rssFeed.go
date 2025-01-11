package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/src/schema"
	"github.com/nathanberry97/rss2go/src/services"
)

func postRssFeed(router *gin.Engine) {
	router.POST("/rss_feed", func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		var rssPostBody schema.RssPostBody
		err = json.Unmarshal(body, &rssPostBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		if rssPostBody.URL == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
			return
		}

		dbConn := services.DatabaseConnection()
		defer dbConn.Close(context.Background())

		id, err := services.PostRssFeed(dbConn, rssPostBody)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error posting RSS feed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "RSS feed posted successfully", "id": id})
	})
}

func getRssFeeds(router *gin.Engine) {
	router.GET("/rss_feed", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feeds retrieved successfully"})
	})
}

func deleteRssFeed(router *gin.Engine) {
	router.DELETE("/rss_feed/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Rss feed deleted successfully"})
	})
}
