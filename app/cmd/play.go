package cmd

import (
	"gohub/pkg/console"
	"gohub/pkg/redis"
	"time"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
	Args:  cobra.NoArgs,
}

func runPlay(cmd *cobra.Command, args []string) {
	// 存进 redis 中
	redis.Redis.Set("foo", "bar", 100*time.Second)
	// 从 redis 中取出来
	console.Success(redis.Redis.Get("foo"))
}
