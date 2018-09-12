// Author: Xu Fei
// Date: 2018/9/12
package service

import (
	"github.com/go-redis/redis"
	"github.com/xfstart07/watcher/config"
)

var Redis *redis.Client

func ConnRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: config.Config.RedisURI,
		Password: config.Config.RedisPass,
	})

	zlog.Sugar().Info(Redis.Ping().Result())
}
