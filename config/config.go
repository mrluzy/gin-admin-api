package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// config 总配置文件
type config struct {
	System system
	Logger logger
}

// 系统配置
type system struct {
	Host string
	Port int
	Env  string
}

// 日志配置
type logger struct {
	Level        string
	Prefix       string
	director     string
	showLine     bool
	LogInConsole bool
}

var Config *config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("........")
		return
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("----------")
		return
	}
}
