package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vue3-bashItem/pkg/logger"
)

const (
	successCode     = 200
	authErrorCode   = 1001
	permErrorCode   = 1002
	paramsErrorCode = 1003
	logicErrorCode  = 1004
	baseErrorCode   = 1005
	unKnowErrorCode = 1006
)

type responseVue3 struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func Vue3Response(c *gin.Context, data interface{}) {
	ret := &responseVue3{
		Code: "00000",
		Data: data,
		//Msg:  "一切ok",
	}
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是200
}

type responseStruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func Success(c *gin.Context, data interface{}) {
	ret := &responseStruct{
		Code: successCode,
		Data: data,
	}
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是200
}

func AuthError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "认证错误!"
	}
	ret := &responseStruct{
		Code: authErrorCode,
		Msg:  msg,
	}
	logger.FileLogger.Error(msg)
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是10001
}

func PermError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "无权限!"
	}
	ret := &responseStruct{
		Code: permErrorCode,
		Msg:  msg,
	}
	logger.FileLogger.Warn(msg)
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是10002
}

func ParamsError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "参数错误!"
	}
	ret := &responseStruct{
		Code: paramsErrorCode,
		Msg:  msg,
	}
	logger.FileLogger.Warn(msg)
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是10003
}

func LogicError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "逻辑错误!"
	}
	ret := &responseStruct{
		Code: logicErrorCode,
		Msg:  msg,
	}
	logger.FileLogger.Error(msg)
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是10004
}

func BaseError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "基础错误!"
	}
	ret := &responseStruct{
		Code: baseErrorCode,
		Msg:  msg,
	}
	logger.FileLogger.Error(msg)
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是10005
	//c.JSON(http.StatusOK, ret)
}

func UnKnowError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "未知错误!"
	}
	ret := &responseStruct{
		Code: unKnowErrorCode,
		Msg:  msg,
	}
	logger.FileLogger.Error(msg)
	c.JSON(http.StatusOK, ret) // http状态是200 前端自定义判断code是10006
	//c.JSON(http.StatusInternalServerError, ret)
}

// 留下来真正返回异常状态 并返回堆栈信息的func
func ServiceCheckError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "服务检查失败!"
	}
	ret := &responseStruct{
		Code: unKnowErrorCode,
		Msg:  msg,
	}
	//msg = msg + "\n\t" + string(debug.Stack()) // debug.Stack() 堆栈信息
	logger.FileLogger.Error(msg)
	c.JSON(http.StatusInternalServerError, ret)
}
