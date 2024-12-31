package cmd

import (
	"gohub/database/migrations"
	"gohub/pkg/migrate"

	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateRollback = &cobra.Command{
	Use: "down",
	// 设置别名 migrate down == migrate rollback
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	// 初始化迁移器
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	// 获取迁移器
	migrator := migrator()

	// 执行迁移
	migrator.Up()
}

func runDown(cmd *cobra.Command, args []string) {
	// 获取迁移器
	migrator := migrator()

	// 回滚迁移
	migrator.Rollback()
}
