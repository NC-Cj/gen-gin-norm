package v1

import "github.com/gin-gonic/gin"

type MyService struct{}

func NewMyService() *MyService {
	return &MyService{}
}

func (s *MyService) QueryUser(ctx *gin.Context) int {
	return 123
}
