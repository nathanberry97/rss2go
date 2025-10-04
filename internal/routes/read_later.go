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

func postReadLater(router *gin.Engine) {
	router.POST("/partials/later/:id", func(c *gin.Context) {
		articleId := c.Param("id")

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		err := services.PostReadLater(dbConn, articleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error posting Read Later article: " + err.Error()})
			return
		}

		response, _ := components.GenerateArticleButton(`/partials/later/`+articleId, "Read Later", `readlater_`+articleId, true)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}

func deleteReadLater(router *gin.Engine) {
	router.DELETE("/partials/later/:id", func(c *gin.Context) {
		articleId := c.Param("id")

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		err := services.DeleteReadLater(dbConn, articleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Read Later article: " + err.Error()})
			return
		}

		response, _ := components.GenerateArticleButton(`/partials/later/`+articleId, "Read Later", `readlater_`+articleId, false)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}

func getReadLater(router *gin.Engine) {
	router.GET("/partials/later", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
		if err != nil || page < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
		if err != nil || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		articles, err := services.GetReadLater(dbConn, page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting Read Later article: " + err.Error()})
			return
		}

		response, _ := components.GenerateArticleList(articles, nil, schema.ArticlesReadLater)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}
