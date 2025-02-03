package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
)

func feedsPage(router *gin.Engine) {
	router.GET("/feeds", func(c *gin.Context) {
		title := "Feeds"
		formHTML := components.GenerateForm("/partials/feed", "RSS Feed URL")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()

		c.HTML(200, "feeds.html", gin.H{
			"title":    title,
			"form":     formHTML,
			"navbar":   navbar,
			"metadata": metadata,
		})
	})
}

func articlesPage(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		title := "Articles"
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()

		c.HTML(200, "articles.html", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
		})
	})
}
