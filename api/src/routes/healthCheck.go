package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanberry97/rss2go/src/services"
)

func healthCheck(router *gin.Engine) {
	router.GET("/health-check", func(ctx *gin.Context) {
		message, status := "ok", http.StatusOK

		dbConn := services.DatabaseConnection()
		defer dbConn.Close(context.Background())

		err := dbConn.Ping(context.Background())
		if err != nil {
			log.Printf("Query failed: %v\n", err)
			message, status = "Error", http.StatusInternalServerError
		}

		ctx.JSON(status, gin.H{"message": message})
	})
}
