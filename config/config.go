package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	
	c := Config{
		Name: cfg,
	}
	
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	
	// 监控配置文件的变化并热加载
	c.watchConfig()
	
	
	
	return nil
}


// 初始化配置
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 指定了配置文件路径，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 否则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	
	// 设置读取环境变量中的配置
	viper.AutomaticEnv()
	// 需要读取的环境变量的前缀
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	
	// 解析所有配置信息
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	
	return nil
}

// 监控配置文件的变动，实时更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
	
}
