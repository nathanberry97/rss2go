package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/pkg/components"
	"github.com/nathanberry97/rss2go/pkg/services"
)

func feedsPage(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		formHTML := components.GenerateForm("/rss_feeds", "RSS Feed URL")

		conn := database.DatabaseConnection()

		feeds, err := services.GetRssFeeds(conn)
		if err != nil {
			panic(err)
		}

		listHtml := components.GenerateFeedList(feeds)

		c.HTML(200, "index.html", gin.H{
			"title":    "RSS Feeds",
			"form":     formHTML,
			"feedList": listHtml,
		})
	})
}
