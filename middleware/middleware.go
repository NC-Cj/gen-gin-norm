package middleware

import (
	"github.com/NC-Cj/gen-gin-norm/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CrossDomain(context *gin.Context) {
	method := context.Request.Method
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	context.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.Next()
}
func RequestHeaders(context *gin.Context) {
	context.Header("X-Request-ID", pkg.GenerateRequestID())
	context.Header("X-Request-Time", pkg.GenerateRequestTime())
	context.Next()
}
