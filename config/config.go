// Author: Xu Fei
// Date: 2018/9/12
package config

import (
	"github.com/go-ini/ini"
	"log"
)

var Config = Conf{}

type Conf struct {
	RedisURI string `ini:"redis_uri"`
	RedisPass string `ini:"redis_pass"`
	RedisDB int `ini:"redis_db"`
	WatchPaths []string `ini:"watch_path"`
}

func Load() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}

	err = cfg.MapTo(&Config)
	if err != nil {
		log.Fatalln(err)
	}
}