package middleware

import (
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		//logger.FileLogger.Info(fmt.Sprintf("Status:%v Method:%v URL:%v",
		//	c.Writer.Status(), c.Request.Method, c.Request.URL))
		// 如下1行输出用户请求
		logger.FileLogger.Info(fmt.Sprintf("Status====%v Method====%v URL====%v query===%v",
			c.Writer.Status(), c.Request.Method, c.Request.URL, c.Request.URL.RawQuery))

		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = utils.GetRandomUUID()
		}
		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()

	}
}
