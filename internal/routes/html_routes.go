package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
	"github.com/nathanberry97/rss2go/internal/schema"
)

func feedsPage(router *gin.Engine) {
	router.GET("/feeds", func(c *gin.Context) {
		title := "Feeds"
		formHTML := components.GenerateInputForm("/partials/feed", "RSS Feed URL")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()

		c.HTML(200, "feeds.tmpl", gin.H{
			"title":    title,
			"form":     formHTML,
			"navbar":   navbar,
			"metadata": metadata,
		})
	})
}

func articlesPage(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		title := "Latest"
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()
		query := components.GenerateArticleQuery(schema.Articles, nil)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func articlesByFeedPage(router *gin.Engine) {
	router.GET("/articles/:feedId", func(c *gin.Context) {
		feedId := c.Param("feedId")
		title := c.DefaultQuery("title", "Feed")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()
		query := components.GenerateArticleQuery(schema.ArticlesByFeed, &feedId)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func articlesFavourite(router *gin.Engine) {
	router.GET("/articles/favourites", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Favourites")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()
		query := components.GenerateArticleQuery(schema.ArticlesFavourite, nil)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func articlesReadLater(router *gin.Engine) {
	router.GET("/articles/later", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Read Later")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()
		query := components.GenerateArticleQuery(schema.ArticlesReadLater, nil)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func settings(router *gin.Engine) {
	router.GET("/settings", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Settings")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData()

		c.HTML(200, "placeholder.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
		})
	})
}
