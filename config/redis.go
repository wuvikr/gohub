package config

import "gohub/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "192.168.1.191"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"username": config.Env("REDIS_USERNAME", "wuvikr"),
			"password": config.Env("REDIS_PASSWORD", "Admin@098"),

			// 业务类存储使用 1 (图片验证码、短信验证码、会话)
			"database": config.Env("REDIS_MAIN_DB", 1),
		}
	})
}
