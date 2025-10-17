package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/internal/components"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/services"
)

func getArticles(router *gin.Engine) {
	router.GET("/partials/articles", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
		if err != nil || page < 0 {
			c.String(http.StatusBadRequest, "Error: Invalid page parameter")
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
		if err != nil || limit <= 0 {
			c.String(http.StatusBadRequest, "Error: Invalid limit parameter")
			return
		}

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.String(http.StatusInternalServerError, "Error: Database connection failed")
			return
		}

		articles, err := services.GetArticles(dbConn, page, limit)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		response, _ := components.GenerateArticleList(articles, nil, schema.Articles)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}

func getArticlesByFeedId(router *gin.Engine) {
	router.GET("/partials/articles/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
		if err != nil || page < 0 {
			c.String(http.StatusBadRequest, "Error: Invalid page parameter")
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
		if err != nil || limit <= 0 {
			c.String(http.StatusBadRequest, "Error: Invalid limit parameter")
			return
		}

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.String(http.StatusInternalServerError, "Error: Database connection failed")
			return
		}

		articles, err := services.GetArticlesByFeedId(dbConn, page, limit, id)
		if err != nil {
			fmt.Printf("Error=%s", err.Error())
			c.String(http.StatusInternalServerError, "Error: Failed to fetch articles by ID")
			return
		}

		response, _ := components.GenerateArticleList(articles, &id, schema.ArticlesByFeed)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}
