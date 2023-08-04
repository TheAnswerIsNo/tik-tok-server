package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"tik-tok-server/global"
)

// InitializeConfig 初始化参数方法
func InitializeConfig() *viper.Viper {
	config := "config.yaml"

	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	v := viper.New()
	//加载配置文件
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config error: %s \n", err))
	}

	//监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changhed:", in.Name)
		//重新加载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	//赋值
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v

}
