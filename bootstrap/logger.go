package bootstrap

import (
	"gohub/pkg/config"
	"gohub/pkg/logger"
)

func SetupLogger() {
	logger.InitLogger(
		config.GetString("logger.filename"),
		config.GetInt("logger.max_size"),
		config.GetInt("logger.max_backup"),
		config.GetInt("logger.max_age"),
		config.GetBool("logger.compress"),
		config.GetString("logger.log_type"),
		config.GetString("logger.level"),
	)
}
