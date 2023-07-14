package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Setting struct {
	vp *viper.Viper
}

// 监控并重新读取配置文件
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSections()
		})
	}()
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	//  viper.UnmarshalKey(“groups”, &GroupsSetting) 将配置文件中 groups 下的全部配置加载到 GroupsSetting里
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSections() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// viper 读取配置文件
func NewSetting(confPath string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile(confPath)
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{
		vp: vp,
	}
	s.WatchSettingChange()
	return s, nil
}

func Setup(confPath string) {
	//  读取yaml文件 并赋值
	setting, err := NewSetting(confPath)
	if err != nil {
		log.Fatalf("Init config file error: %v", err)
	}

	// 定义变量用于获取yaml配置 一对一关系 name和yaml顶级名字一致 否则探活会报错
	var settingInfo = []map[string]interface{}{
		{"name": "Server", "data": &ServerSetting},
		{"name": "Database", "data": &DatabaseSetting},
		{"name": "Redis", "data": &RedisSetting},
		{"name": "LogConf", "data": &LogConfSetting},
		{"name": "Redis", "data": &RedisSetting},
		{"name": "App", "data": &AppSetting},
		{"name": "Check", "data": &LogConfSetting},
		{"name": "JwtConf", "data": &JwtConfSetting},
		{"name": "DefaultRoute", "data": &DefaultRouteSetting},
		// name和yaml顶级名字一致 否则探活会报错
	}

	for info_i := 0; info_i < len(settingInfo); info_i++ {
		// setting.ReadSection方法中 settingInfo[info_i]["name"] 需要强转译成string
		err = setting.ReadSection(settingInfo[info_i]["name"].(string), settingInfo[info_i]["data"])
		if err != nil {
			log.Fatalf("%v获取配置配置失败: %v", settingInfo[info_i]["name"], err)
		}
	}

}
