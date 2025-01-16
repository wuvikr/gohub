package cache

import (
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(address, username, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.NewClient(address, username, password, db)
	rs.KeyPrefix = config.GetString("app.name") + ":cache:"
	return rs
}

func (s RedisStore) Set(key string, value string, expiration time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expiration)
}

func (s RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

func (s RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}

func (s RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}

func (s RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}

func (s RedisStore) Increment(key string, value ...int64) {
	s.RedisClient.Increment(key, value...)
}

func (s RedisStore) Decrement(key string, value ...int64) {
	s.RedisClient.Decrement(key, value...)
}
