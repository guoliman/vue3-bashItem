package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"vue3-bashItem/crond"
	_ "vue3-bashItem/docs" // 项目yiyongTest  docs是目录
	"vue3-bashItem/model"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/migrate"
	"vue3-bashItem/pkg/portDetection"
	"vue3-bashItem/pkg/redis"
	"vue3-bashItem/pkg/settings"
	"vue3-bashItem/router"
)

var mode string
var confPath string

// 项目运行前的基础方法
func Setup() {
	settings.Setup(confPath) // 配置文件
	logger.Setup()           // 日志
	portDetection.Setup()    // 探活配置
	model.Setup()            // 数据库
	redis.Setup()            // redis
}

func runServer() {
	engine := router.InitRouter() //路由

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.ServerSetting.Port),
		Handler:        engine,
		ReadTimeout:    settings.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   settings.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//logger.Logger.Info("listen",
		//	zap.String("Host", "*"),
		//	zap.Int("Port", settings.ServerSetting.Port),
		//	)
		listenInfo := fmt.Sprintf("Service start Success 0.0.0.0:%v", settings.ServerSetting.Port)
		logger.Logger.Info(listenInfo)
		logger.FileLogger.Info(listenInfo)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errData := fmt.Sprintf("Server Start Error: %v", zap.Error(err))
			logger.FileLogger.Error(errData)
			panic(errData)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Logger.Info("Shutdown Server ... ")
	logger.FileLogger.Info("Shutdown Server ... ")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) //N秒后停止服务
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		errData := fmt.Sprintf("Server Shutdown Error: %v", zap.Error(err))
		logger.FileLogger.Error(errData)
		logger.Logger.Error(errData)
		os.Exit(1)
	}
	logger.Logger.Info("Server Shutdown.")
	logger.FileLogger.Info("Server Shutdown.")
}

// =========================swagger页面展现信息 （不包含接口信息）方式一 ===========================
// @title     		lplsoc Example Api
// @version   		v1.0
// @host      		localhost:8095
// @BasePath  		/api
// @description     lplsoc是go基础模板

// Apache版本
// @license.name  Apache 2.01
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// 可选多种请求方式
//// @schemes http https
func main() {
	var itemConfigFile string

	sysVersion := runtime.GOOS // 获取系统类型
	if sysVersion == "linux" {
		itemConfigFile = "./configs/prodConfig.yaml"
	} else {
		itemConfigFile = "./configs/testConfig.yaml"
	}

	// 接收参数
	flag.StringVar(&mode, "mode", "server", "运行模式,运行server服务或者执行migrate.")
	flag.StringVar(&confPath, "c", itemConfigFile, "配置文件路径.")
	flag.Parse()

	// 初始化数据库(默认配置文件)  go run main.go -mode=migrate 或者 go run main.go -mode migrate
	// 初始化数据库 指定配置文件  ./vue3-bashItem -c configs/prodConfig.yaml -mode=migrate

	// 启动服务(默认配置文件)     go run main.go
	// 启动服务(指定配置文件)     go run main.go -c=/Users/aa/vue3-bashItem/configs/config.yaml

	if mode != "server" && mode != "migrate" {
		log.Fatal("run 参数非法: 运行模式只能选择 server 或 migrate") // FAtal 服务不会继续执行了
	}

	Setup() // 初始化

	// 根据参数执行对应模式服务
	if mode == "server" {
		crond.RunTask() // 定时任务
		runServer()     // 启动服务
	} else {
		migrate.MigrateTable() // 数据库初始化
	}
}

/*
swag-web  http://localhost:8095/api/swagger/index.html
*/
