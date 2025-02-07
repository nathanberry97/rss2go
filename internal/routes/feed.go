package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/services"
)

func postFeed(router *gin.Engine) {
	router.POST("/partials/feed", func(c *gin.Context) {
		var rssPostBody schema.RssPostBody
		if err := c.ShouldBind(&rssPostBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data: " + err.Error()})
			return
		}

		if rssPostBody.URL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
			return
		}

		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		err := services.PostFeed(dbConn, rssPostBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error posting RSS feed: " + err.Error()})
			return
		}

		c.Header("HX-Trigger", "refreshFeed")
		c.Status(http.StatusOK)
	})
}

func getFeeds(router *gin.Engine) {
	router.GET("/partials/feed", func(c *gin.Context) {
		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		feeds, err := services.GetFeeds(dbConn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving RSS feeds: " + err.Error()})
			return
		}

		updatedFeedListHtml := components.GenerateFeedList(feeds)

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(updatedFeedListHtml))
	})
}

func deleteFeed(router *gin.Engine) {
	router.DELETE("/partials/feed/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format: " + err.Error()})
			return
		}

		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		err = services.DeleteFeed(dbConn, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting RSS feed: " + err.Error()})
			return
		}

		c.Header("HX-Trigger", "refreshFeed")
		c.Status(http.StatusOK)
	})
}
