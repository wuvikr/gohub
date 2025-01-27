package bootstrap

import (
	"fmt"
	"gohub/pkg/cache"
	"gohub/pkg/config"
)

func SetupCache() {

	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)
	cache.InitWithCacheStore(rds)
}
