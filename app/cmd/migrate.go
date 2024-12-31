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

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all migrations",
	Run:   runReset,
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Rollback all migrations and re-run them",
	Run:   runRefresh,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateReset,
		CmdMigrateRefresh,
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

func runReset(cmd *cobra.Command, args []string) {
	// 获取迁移器
	migrator := migrator()

	// 重置迁移
	migrator.Reset()
}

func runRefresh(cmd *cobra.Command, args []string) {
	// 获取迁移器
	migrator := migrator()

	// 刷新迁移
	migrator.Refresh()
}
