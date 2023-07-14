package testExample

import (
	"vue3-bashItem/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* ============如下 写在任意目录都行 不需要引用会自动引用=============
接口参数 https://topgoer.cn/docs/swaggo/swaggo-1diseukku0gfd
// @Summary						 // 接口说明
// @Description					 // 接口详情介绍
// @Accept multipart/form-data   // 获取数据类型
// @Produce application/json     // API可以生成的MIME类型的列表
// @Param dataInfo body string true "post传参"  // 参数传参 这是post传参 定义一个传全部参数
// @Param stu_id formData string true "学号"    // 参数传参 这是get传参 定义一个传1个参数 true必须传参 false非必填
// @Success 200 {string}  string  "aaa--"// 返回数据类型
// @Success 200 {json} json ""msg":"登录成功","code":200,"data":{}"  // 返回json类型 不好使
// @Success 200 {object}  model.Account // 自定义返回类型 model.Account只写模块不写路径 例 /aa/bb/model/modeA.go下的Account 只写model.Account
// @Failure 400 {object}  HTTPError  // 500报错定义
// @Failure 404 "未找到此用户"     // 404报错定义
// @Tags        				// 组名 类似a组下多个接口
// @Router /login [POST]         // 路由地址 这是手写的 需要和真正的路由匹配 /login是路由 post是类型
===============================================================
*/

// @Summary  登录接口 ---接口说明
// @Description 用于用户认证--接口详情介绍
// @Accept multipart/form-data
// @Produce application/json
// @Param stu_id formData string true "学号"
// @Param password formData string false "密码"
// @Success 200 {object}  testExampleConf.HTTPSuccess
// @Failure 400 {object}  testExampleConf.HTTPError
// @Failure 500 "未找到此用户"
// @Tags        bb
// @Router /swag/loginAA [POST]
func LoginBB(c *gin.Context) {
	//logger.Logger.Debug("update 新创建,不存在")     // 命令行输出
	logger.FileLogger.Debug("Login=====") // 文件输出
	c.JSON(200, gin.H{
		"msg":  "找到此用户",
		"code": 200,
	})
}

// @Summary 	bb接口
// @Description bb接口详情介绍
// @Accept      json
// @Produce     json
// @Success 200 {object}  testExampleConf.Response
// @Failure 400 {object}  testExampleConf.HTTPError
// @Tags        gavin
// @Router      /swag/aa/bb [get]
func Helloworld(g *gin.Context) {
	logger.FileLogger.Debug("Helloworld=====")
	g.JSON(http.StatusOK, "helloworld------")
}

// @Summary 	cc接口
// @Description cc接口详情介绍
// @Accept      json
// @Produce     json
// @Success     200 {string}  string  "answer--"
// @Failure 400 {object}  testExampleConf.HTTPError
// @Failure 404 "404报错"
// @Failure 500 "500报错"
// @Tags        gavin
// @Router      /swag/aa/cc [get]
func Bb(g *gin.Context) {
	logger.FileLogger.Info("Bb=====")
	g.JSON(http.StatusOK, "Bb------")
}
