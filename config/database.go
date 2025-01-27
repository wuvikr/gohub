package config

import "gohub/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认数据库
			"connection": config.Env("DB_CONNECTION", "mysql"),

			// 数据库连接信息
			"mysql": map[string]interface{}{
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gohub"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",

				// 链接池配置
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_sseconds":    config.Env("DB_MAX_LIFE_SECONDS", 60*5),
			},

			"sqlite": map[string]interface{}{
				"database": config.Env("DB_SQL_FILE", "database/database.db"),
			},
		}
	})
}
