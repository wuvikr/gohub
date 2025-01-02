package seeders

import (
	"fmt"
	"gohub/database/factories"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		// 创建 10 个用户
		users := factories.MakeUsers(10)

		// 批量创建用户
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
