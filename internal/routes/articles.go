package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/services"
)

func getArticles(router *gin.Engine) {
	router.GET("/partials/articles", func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "0")
		limitStr := c.DefaultQuery("limit", "50")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		articles, err := services.GetArticles(dbConn, page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := components.GenerateArticleList(articles)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}
