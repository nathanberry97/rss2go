package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(router *gin.Engine) {
	router.GET("/health-check", func(ctx *gin.Context) {
		message, status := "ok", http.StatusOK

		ctx.JSON(status, gin.H{"message": message})
	})
}
