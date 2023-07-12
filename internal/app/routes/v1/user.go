package v1

import (
	"github.com/NC-Cj/gen-gin-norm/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, c *gin.Context) {

	r.GET("/user", controllers.Controller.QueryUser())
}
