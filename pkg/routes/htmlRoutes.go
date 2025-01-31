package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
)

func feedsPage(router *gin.Engine) {
	router.GET("/feeds", func(c *gin.Context) {
		formHTML := components.GenerateForm("/partials/feed", "RSS Feed URL")
		navbar := components.GenerateNavbar()

		c.HTML(200, "feeds.html", gin.H{
			"title":  "RSS Feeds",
			"form":   formHTML,
			"navbar": navbar,
		})
	})
}

func articlesPage(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		navbar := components.GenerateNavbar()

		c.HTML(200, "articles.html", gin.H{
			"title":  "RSS Articles",
			"navbar": navbar,
		})
	})
}
