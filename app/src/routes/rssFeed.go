package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/src/schema"
	"github.com/nathanberry97/rss2go/src/services"
)

// postRssFeed handles POST requests to add a new RSS feed
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
		defer dbConn.Close()

		id, err := services.PostRssFeed(dbConn, rssPostBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error posting RSS feed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "RSS feed posted successfully", "id": id})
	})
}

// getRssFeeds handles GET requests to retrieve all RSS feeds
func getRssFeeds(router *gin.Engine) {
	router.GET("/rss_feed", func(ctx *gin.Context) {
		dbConn := services.DatabaseConnection()
		defer dbConn.Close()

		feeds, err := services.GetRssFeeds(dbConn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving RSS feeds"})
			return
		}

		ctx.JSON(http.StatusOK, feeds)
	})
}

func deleteRssFeed(router *gin.Engine) {
	router.DELETE("/rss_feed/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		dbConn := services.DatabaseConnection()
		defer dbConn.Close()

		err = services.DeleteRssFeed(dbConn, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting RSS feed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "RSS feed deleted successfully"})
	})
}
