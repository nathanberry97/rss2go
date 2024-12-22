package routes

import "github.com/gin-gonic/gin"

func InitialiseRouter() *gin.Engine {
	router := gin.Default()
	healthCheck(router)
	return router
}
