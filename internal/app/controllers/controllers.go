package controllers

import (
	"github.com/NC-Cj/gen-gin-norm/internal/app/logic/v1"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *v1.MyService
}

func NewController(s *v1.MyService) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) QueryUser(ctx *gin.Context) int {
	return c.service.QueryUser(ctx)
}
