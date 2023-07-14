package portDetection

import (
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/settings"
	"fmt"
	"net"
	"time"
)

func CheckPorts(serverName string, ip_port string) {
	// 检测端口
	conn, err := net.DialTimeout("tcp", ip_port, 3*time.Second) // 3秒超时
	if err != nil {
		errorData := fmt.Sprintf("探活失败 %s %s 第一类报错: %s", serverName, ip_port, err)
		logger.FileLogger.Error(errorData)
		panic(errorData) // 主动异常退出
	} else {
		if conn != nil {
			logger.Logger.Info("探活成功")
			logger.FileLogger.Info("探活成功")
			conn.Close()
		} else {
			errorData := fmt.Sprintf("探活失败 %s %s 第二类报错: %s", serverName, ip_port, conn)
			logger.FileLogger.Error(errorData)
			panic(errorData) // 主动异常退出
		}
	}
}

func Setup() {
	//CheckPorts("worktile.1yongcloud.com:80")
	//CheckPorts("192.168.148.198:3306")

	// mysql 探活
	CheckPorts("mysql", fmt.Sprintf("%v:%v", settings.DatabaseSetting.Host, settings.DatabaseSetting.Port))

	// redis 探活
	CheckPorts("redis", fmt.Sprintf("%v", settings.RedisSetting.Host))

}
