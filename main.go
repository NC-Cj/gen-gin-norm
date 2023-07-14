package main

import (
	"github.com/NC-Cj/gen-gin-norm/core"
	"github.com/NC-Cj/gen-gin-norm/internal/app"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.Default()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: nil,
	}))
	r.Use(core.RequestHandler())

	app.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
