package app

import (
	"gohub/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, err := time.LoadLocation(config.Get("app.timezone"))
	if err != nil {
		// 默认为本地时区
		return time.Now()
	}
	return time.Now().In(chinaTimezone)
}

func URL(path string) string {
	return config.Get("app.url") + path
}

func V1URL(path string) string {
	return URL("/v1/") + path
}
