package redis

import (
	"context"
	"gohub/pkg/logger"
	"sync"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

var (
	once  sync.Once    // once 确保全局的 Redis 对象只实例一次
	Redis *RedisClient // 全局Redis实例
)

func ConnectRedis(address, username, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

func NewClient(address, username, password string, db int) *RedisClient {

	// 初始化实例
	rds := &RedisClient{}
	rds.Context = context.Background()

	// 连接数据库
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试连接
	err := rds.Client.Ping(rds.Context).Err()
	logger.LogIf(err)

	return rds
}

// Ping 测试连接
func (rds *RedisClient) Ping() error {
	return rds.Client.Ping(rds.Context).Err()
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds *RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds *RedisClient) Get(key string) string {
	val, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return val
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false

func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds *RedisClient) Del(key ...string) bool {
	if err := rds.Client.Del(rds.Context, key...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空当前数据库
func (rds *RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Close 关闭 Redis 连接
func (rds *RedisClient) Close() {
	_ = rds.Client.Close()
}

// Increment 增加, 支持两个参数，第一个为 key，第二个为增加的值（可选）
// 如果只有一个参数，则为增加 1, 如果有第二个参数，则为增加该值
func (rds *RedisClient) Increment(key string, value ...int64) bool {
	switch len(value) {
	case 0:
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Incr", err.Error())
			return false
		}
	case 1:
		if err := rds.Client.IncrBy(rds.Context, key, value[0]).Err(); err != nil {
			logger.ErrorString("Redis", "IncrBy", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}
	return true
}

// Decrement 减少, 支持两个参数，第一个为 key，第二个为减少的值（可选）
// 如果第二个参数没有，默认减去 1，如果有则为减少该值
func (rds *RedisClient) Decrement(key string, value ...int64) bool {
	switch len(value) {
	case 0:
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decr", err.Error())
			return false
		}
	case 1:
		if err := rds.Client.DecrBy(rds.Context, key, value[0]).Err(); err != nil {
			logger.ErrorString("Redis", "DecrBy", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false

	}
	return true
}
