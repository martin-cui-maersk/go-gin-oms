package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/martin-cui-maersk/go-gin-oms/config"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"github.com/spf13/viper"
	"os"
)

// Viper 配置
func Viper() *viper.Viper {
	// 获取config file
	configFile := getConfigFilePath()
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error to read config file, %w", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Server); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Server); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
	return v
}

// getConfigFilePath 获取配置文件路径
func getConfigFilePath() (path string) {
	// 设置运行模式 默认debug
	var mode string
	flag.StringVar(&mode, "m", "", "set run mode.")
	flag.Parse()
	switch mode {
	case gin.DebugMode:
		path = config.ConfigDebugFile
	case gin.ReleaseMode:
		path = config.ConfigReleaseFile
	case gin.TestMode:
		path = config.ConfigTestFile
	default:
		mode = gin.DebugMode
		path = config.ConfigDebugFile
	}
	gin.SetMode(mode)
	fmt.Printf("您正在使用 gin 的 %s 模式运行, 配置文件为 %s\n", gin.Mode(), path)
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		path = config.ConfigDebugFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", path)
	}
	return
}
