package main

import (
	"gin-admin-api/config"
	"gin-admin-api/core"
	"gin-admin-api/global"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化日志
	global.Log = core.InitLogger()

	// 初始化mysql
	core.InitMysql()

	// 初始化 redis
	core.InitRedis()
}
