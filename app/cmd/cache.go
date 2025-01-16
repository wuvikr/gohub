package cmd

import (
	"fmt"
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

var CmdCacheForgot = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cache-key",
	Run:   runCacheForgot,
}

var key string

func init() {
	CmdCacheForgot.Flags().StringVarP(&key, "key", "k", "", "需要清理的缓存 key 名称")

	CmdCache.AddCommand(
		CmdCacheClear,
		CmdCacheForgot,
	)
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("清除缓存成功")
}

func runCacheForgot(cmd *cobra.Command, args []string) {
	// 删除指定缓存
	cache.Forget(key)
	message := fmt.Sprintf("缓存 %s 已清除", key)
	console.Success(message)
}
