package core

import (
	"gin-admin-api/config"
	"gin-admin-api/global"
	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

func InitRedis() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Db})

	_, err := RedisDb.Ping(global.Ctx).Result()
	if err != nil {
		global.Log.Errorf("[redis] init fail")
		return err
	}
	global.Log.Infof("[redis] init success")
	return nil
}
