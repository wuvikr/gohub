package cmd

import (
	"gohub/pkg/cache"
	"gohub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

func init() {
	CmdCache.AddCommand(
		CmdCacheClear,
	)
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("清除缓存成功")
}
