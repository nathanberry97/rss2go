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
		formHTML := components.GenerateFeedInputForm("/partials/feed", "RSS Feed URL")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData(cssFile)

		c.HTML(200, "feed.tmpl", gin.H{
			"title":    title,
			"form":     formHTML,
			"navbar":   navbar,
			"metadata": metadata,
		})
	})
}

func articlesPage(router *gin.Engine, cssFile string) {
	router.GET("/", func(c *gin.Context) {
		title := "Latest"
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData(cssFile)
		query := components.GenerateArticleQuery(schema.Articles, nil)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
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

		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData(cssFile)
		query := components.GenerateArticleQuery(schema.ArticlesByFeed, &feedId)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func articlesFavourite(router *gin.Engine, cssFile string) {
	router.GET("/articles/favourites", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Favourites")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData(cssFile)
		query := components.GenerateArticleQuery(schema.ArticlesFavourite, nil)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func articlesReadLater(router *gin.Engine, cssFile string) {
	router.GET("/articles/later", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Read Later")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData(cssFile)
		query := components.GenerateArticleQuery(schema.ArticlesReadLater, nil)

		c.HTML(200, "articles.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
			"query":    query,
		})
	})
}

func settings(router *gin.Engine, cssFile string) {
	router.GET("/settings", func(c *gin.Context) {
		title := c.DefaultQuery("title", "Settings")
		navbar := components.GenerateNavbar()
		metadata := components.GenerateMetaData(cssFile)

		c.HTML(200, "placeholder.tmpl", gin.H{
			"title":    title,
			"navbar":   navbar,
			"metadata": metadata,
		})
	})
}
