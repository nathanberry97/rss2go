package routes

import (
	"io"
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
		c.Status(http.StatusNoContent)
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

		updatedFeedListHtml, _ := components.GenerateFeedList(feeds)

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
		c.Status(http.StatusNoContent)
	})
}

func postFeedOpml(router *gin.Engine) {
	router.POST("/partials/feed/opml", func(c *gin.Context) {
		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		fileHeader, err := c.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error cannot find file: " + err.Error()})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error cannot open file: " + err.Error()})
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error cannot read content of file: " + err.Error()})
			return
		}

		err = services.PostFeedOpml(dbConn, content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error cannot post OPML content: " + err.Error()})
			return
		}

		c.Header("HX-Trigger", "refreshFeed")
		c.Status(http.StatusNoContent)
	})
}

func getFeedOpml(router *gin.Engine) {
	router.GET("/partials/feed/opml", func(c *gin.Context) {
		dbConn := database.DatabaseConnection()
		defer dbConn.Close()

		c.Header("Content-Disposition", `attachment; filename="rss2go.xml"`)
		c.Header("Content-Type", "application/xml")

		xmlContent, err := services.GetFeedsOpml(dbConn)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to generate OPML")
			return
		}

		c.Data(http.StatusOK, "application/xml", xmlContent)
	})
}
