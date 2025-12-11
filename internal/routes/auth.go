package routes

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func basicAuthMiddleware(user, pass string) gin.HandlerFunc {
	expected := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))

	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth != expected {
			ctx.Header("WWW-Authenticate", "Basic")
			ctx.AbortWithStatus(401)
			return
		}
		ctx.Next()
	}
}
