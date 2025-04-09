package middleware

import (
	"fmt"
	"net"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
					logger.FileLogger.Error(fmt.Sprintf("%v %v %v", c.Request.URL.Path),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				// 记录错误日志
				var errMsg string
				if e, ok := err.(error); ok {
					errMsg = e.Error()
				} else {
					errMsg = fmt.Sprintf("%v", err)
				}

				// 获取实际的报错位置
				pc := make([]uintptr, 10)
				n := runtime.Callers(3, pc) // 跳过3层调用栈
				frames := runtime.CallersFrames(pc[:n])
				var panicLocation string
				if frame, more := frames.Next(); more {
					panicLocation = fmt.Sprintf("%s:%d", frame.File, frame.Line)
				}

				// 输出错误信息和报错位置
				fullError := fmt.Sprintf("CodePATH: %s\n    MESSAGE: %v ", panicLocation,errMsg)
				// logger.Logger.Error(fullError)
				// logger.FileLogger.Error(fullError)
				response.UnKnowError(c, fullError)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
