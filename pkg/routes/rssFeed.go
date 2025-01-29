package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/pkg/components"
	"github.com/nathanberry97/rss2go/pkg/schema"
	"github.com/nathanberry97/rss2go/pkg/services"
)

func postRssFeed(router *gin.Engine) {
	router.POST("/rss_feed", func(ctx *gin.Context) {
		var rssPostBody schema.RssPostBody
		if err := ctx.ShouldBind(&rssPostBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data: " + err.Error()})
			return
		}

		fmt.Println(rssPostBody)
		if rssPostBody.URL == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
			return
		}

		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		id, err := services.PostRssFeed(dbConn, rssPostBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error posting RSS feed: " + err.Error()})
			return
		}

		fmt.Println(id)

		ctx.Header("HX-Trigger", "refreshFeed")
		ctx.Status(http.StatusOK)
	})
}

func getRssFeeds(router *gin.Engine) {
	router.GET("/rss_feed", func(ctx *gin.Context) {
		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		feeds, err := services.GetRssFeeds(dbConn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving RSS feeds: " + err.Error()})
			return
		}

		updatedFeedListHtml := components.GenerateFeedList(feeds)

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(updatedFeedListHtml))
	})
}

func deleteRssFeed(router *gin.Engine) {
	router.DELETE("/rss_feed/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format: " + err.Error()})
			return
		}

		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		err = services.DeleteRssFeed(dbConn, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting RSS feed: " + err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "RSS feed deleted successfully"})
	})
}
