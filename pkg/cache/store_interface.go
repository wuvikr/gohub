package cache

import "time"

type Store interface {
	Get(key string) string
	Set(key string, value string, expiration time.Duration)
	Has(key string) bool
	Forget(key string)
	Forever(key string, value string)
	Flush()

	IsAlive() error

	// 当参数只有一个时，参数为 key，增加 1
	// 当参数有两个时，第一个参数为 key，第二个参数为增加的值 int64 类型
	Increment(key string, value ...int64)

	// 当参数只有一个时，参数为 key，减去 1
	// 当参数有两个时，第一个参数为 key，第二个参数为减去的值 int64 类型
	Decrement(key string, value ...int64)
}
