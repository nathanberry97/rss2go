package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheck(router *gin.Engine) {
	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
