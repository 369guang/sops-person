package settings

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct{}

const defaultFile = "application.yaml"

func Init() error {
	c := Config{}
	if err := c.InitConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) InitConfig() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func SetConfig(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "请输入配置文件")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv("CONFIG"); configEnv == "" {
				config = defaultFile
			} else {
				config = defaultFile
			}
		}
	} else {
		config = path[0]
	}

	obj := viper.New()
	obj.AddConfigPath("conf")
	obj.SetConfigFile(config)
	obj.SetConfigType("yaml")
	err := obj.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("错误的配置文件，原因: %v", err))
	}
	obj.WatchConfig()

	obj.OnConfigChange(func(f fsnotify.Event) {
		fmt.Println("配置文件刷新:", f.Name)
	})

	return obj
}
