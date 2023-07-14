package core

import (
	"github.com/NC-Cj/gen-gin-norm/core/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"time"
)

// RequestHandler 是一个中间件，用于处理请求和响应
func RequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置跨域
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		requestID := xid.New().String()
		requestTime := time.Now()

		// 将请求ID和时间添加到Gin的上下文中，这样在处理程序中可以获取它们
		c.Set("traceid", requestID)
		// 调用后面的处理程序
		c.Next()

		logger.Info("%s | %s | %s | %v", c.Request.Method, c.Request.URL, requestID, requestTime)
		// 将请求ID和时间添加到响应头
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Writer.Header().Set("X-Request-Time", requestTime.Format(time.RFC3339))
	}
}
