package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/services"
)

func feedsPage(router *gin.Engine, cssFile string) {
	router.GET("/feeds", func(c *gin.Context) {
		title := "Feeds"
		formHTML, _ := components.GenerateFeedInputForm("/partials/feed", "RSS Feed URL")
		opmlButton, _ := components.GenerateOPMLButton("/partials/feed/opml")
		metadata, _ := components.GenerateMetaData(cssFile)
		navbar, _ := components.GenerateNavbar()

		c.HTML(200, "base.tmpl", gin.H{
			"title":    title,
			"form":     formHTML,
			"feed":     true,
			"navbar":   navbar,
			"metadata": metadata,
			"opml":     opmlButton,
		})
	})
}

func articlesPage(router *gin.Engine, cssFile string) {
	router.GET("/", func(c *gin.Context) {
		title := "Latest"
		metadata, _ := components.GenerateMetaData(cssFile)
		query, _ := components.GenerateArticleQuery(schema.Articles, nil)
		navbar, _ := components.GenerateNavbar()

		c.HTML(200, "base.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"articles": true,
			"query":    query,
		})
	})
}

func articlesByFeedPage(router *gin.Engine, cssFile string) {
	router.GET("/articles/:feedId", func(c *gin.Context) {
		feedId := c.Param("feedId")

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		title, err := services.GetArticleName(dbConn, feedId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get feed name"})
			return
		}

		navbar, _ := components.GenerateNavbar()
		metadata, _ := components.GenerateMetaData(cssFile)
		query, _ := components.GenerateArticleQuery(schema.ArticlesByFeed, &feedId)

		c.HTML(200, "base.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"articles": true,
			"query":    query,
		})
	})
}

func articlesFavourite(router *gin.Engine, cssFile string) {
	router.GET("/articles/favourites", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Favourites")
		navbar, _ := components.GenerateNavbar()
		metadata, _ := components.GenerateMetaData(cssFile)
		query, _ := components.GenerateArticleQuery(schema.ArticlesFavourite, nil)

		c.HTML(200, "base.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"articles": true,
			"query":    query,
		})
	})
}

func articlesReadLater(router *gin.Engine, cssFile string) {
	router.GET("/articles/later", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Read Later")
		navbar, _ := components.GenerateNavbar()
		metadata, _ := components.GenerateMetaData(cssFile)
		query, _ := components.GenerateArticleQuery(schema.ArticlesReadLater, nil)

		c.HTML(200, "base.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"articles": true,
			"query":    query,
		})
	})
}
