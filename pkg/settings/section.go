package settings

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	User        string
	Password    string
	Host        string
	Port        int
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type App struct {
	PageSize               int
	UserCenterGetOwnApiURL string
	SessionCookieName      string
}

type LogConf struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

type Check struct {
	Users []string
}

type JwtConf struct {
	JwtKey  string
	JwtTime int
}

var (
	CheckSetting        *Check
	ServerSetting       *ServerSettingS
	DatabaseSetting     *Database
	RedisSetting        *Redis
	AppSetting          *App
	LogConfSetting      *LogConf
	JwtConfSetting      *JwtConf
	DefaultRouteSetting []map[string]string
)

var sections = make(map[string]interface{})

//var settingInfo = make([]map[string]interface{}, 0)
//var settingInfo []map[string]interface{}
