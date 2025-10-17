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

func postFavourite(router *gin.Engine) {
	router.POST("/partials/favourite/:id", func(c *gin.Context) {
		articleId := c.Param("id")

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.String(http.StatusInternalServerError, "Database connection failed")
			return
		}

		err := services.PostFavourite(dbConn, articleId)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		response, _ := components.GenerateArticleButton(`/partials/favourite/`+articleId, "Favourite", `favourite_`+articleId, true)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}

func deleteFavourite(router *gin.Engine) {
	router.DELETE("/partials/favourite/:id", func(c *gin.Context) {
		articleId := c.Param("id")

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.String(http.StatusInternalServerError, "Database connection failed")
			return
		}

		err := services.DeleteFavourite(dbConn, articleId)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		response, _ := components.GenerateArticleButton(`/partials/favourite/`+articleId, "Favourite", `favourite_`+articleId, false)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}

func getFavourites(router *gin.Engine) {
	router.GET("/partials/favourite", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
		if err != nil || page < 0 {
			c.String(http.StatusBadRequest, "Invalid page parameter")
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
		if err != nil || limit <= 0 {
			c.String(http.StatusBadRequest, "Invalid limit parameter")
			return
		}

		dbConn := database.DatabaseConnection()
		if dbConn == nil {
			c.String(http.StatusInternalServerError, "Database connection failed")
			return
		}

		articles, err := services.GetFavourites(dbConn, page, limit)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		response, _ := components.GenerateArticleList(articles, nil, schema.ArticlesFavourite)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	})
}
