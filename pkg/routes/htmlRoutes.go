package routes

import (
	"github.com/gin-gonic/gin"
)

func homepage(router *gin.Engine) {

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home Page",
		})
	})
}
