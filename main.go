package main

import (
	"fmt"
	"gin-admin-api/config"
)

func main() {
	// 初始化配置
	config.InitConfig()
	fmt.Println("系统配置： ", config.Config.System)
	fmt.Println("日志配置： ", config.Config.Logger)
}
