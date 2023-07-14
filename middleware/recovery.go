package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/response"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.FileLogger.Error(fmt.Sprintf("%v %v %v",
						c.Request.URL.Path),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					//logger.Logger.Error(c.Request.URL.Path,
					//	zap.Any("error", err),
					//	zap.String("request", string(httpRequest)),
					//)
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				// 记录错误日志
				logger.Logger.Error(err.(error).Error())
				logger.Logger.Error(string(debug.Stack()))
				logger.FileLogger.Error(fmt.Sprintf("Recovery error: %v", err.(error).Error()))
				logger.FileLogger.Error(fmt.Sprintf("Recovery error: %v", string(debug.Stack()))) //文件内输出没格式
				response.UnKnowError(c, err.(error).Error())
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
