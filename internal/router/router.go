package router

import (
	"github.com/NC-Cj/gen-gin-norm/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	e := gin.New()
	e.Use(middleware.CrossDomain)
	e.Use(middleware.RequestHeaders)

	return e
}
