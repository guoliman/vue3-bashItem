package middleware

import (
	"fmt"
	"time"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 输出请求头
		// logger.FileLogger.Info(fmt.Sprintf("Request Headers: %v", c.Request.Header))

		// 获取请求ID
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = utils.GetRandomUUID()
		}
		c.Set("X-Request-Id", requestId) // 设置到上下文中
		// c.Writer.Header().Set("X-Request-Id", requestId) // 设置到响应头中  这么设置无效
		c.Request.Header.Set("X-Request-Id", requestId) // 设置到请求头中 这么设置生效
		// logger.FileLogger.Info(fmt.Sprintf("Request X-Request-Id: %v", c.Request.Header.Get("X-Request-Id")))

		// 请求入口时记录请求信息
		logger.FileLogger.Info(fmt.Sprintf("Request -------- X-Request-Id: %v, URL: %v, Type: %v, Query: %v",
			requestId, c.Request.URL, c.Request.Method, c.Request.URL.RawQuery))

		// 接口处理逻辑
		c.Next()

		// 计算请求处理时间  单位是微秒  1000微秒（µs）等于1毫秒（ms）， 1000毫秒（ms）等于1秒（s）。
		requestTime := time.Since(startTime)

		// 请求返回时记录响应信息
		logger.FileLogger.Info(fmt.Sprintf("Response -------- X-Request-Id: %v, URL: %v, Type: %v, Code: %v, RequestTime: %v",
			requestId, c.Request.URL, c.Request.Method, c.Writer.Status(), requestTime))
	}
}
