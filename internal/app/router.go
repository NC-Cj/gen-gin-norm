package app

import (
	"github.com/NC-Cj/gen-gin-norm/internal/app/routes/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 将中间件应用到所有的路由
	r.Use(requestHandler())

	v1.SetupUserRoutes()

	return r
}
