package conf

import (
	"connector/lib"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GatewayCompanyConfigDir string
var ManagerAddr string
var ConnectorServer string
var NgxSharedSetUrl string
var infoLog lib.InfoLogger

// Config struct
type Config struct {
	Name string
}

// Init config
// @Param config_path
// @Return err
func Init(cfg string) error {
	infoLog.Println("Init conf")
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//c.WatchConfig()
	return nil
}

// Init config private
// @Param
// @Return err
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	GatewayCompanyConfigDir = viper.GetString("path.jsonpath")
	ManagerAddr = viper.GetString("path.manager_addr")
	NgxSharedSetUrl = viper.GetString("path.ngx_shared_url")
	watchConfig()
	return nil
}

// WatchConfig hot modify
// 监听配置文件是否改变,用于热更新
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		infoLog.Printf("Config file changed: %s\n", e.Name)
		GatewayCompanyConfigDir = viper.GetString("path.jsonpath")
		ManagerAddr = viper.GetString("path.manager_addr")
		NgxSharedSetUrl = viper.GetString("path.ngx_shared_url")
		lib.ModifyLevel()
	})
}
