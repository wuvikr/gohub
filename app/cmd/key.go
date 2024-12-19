package cmd

import (
	"gohub/pkg/console"
	"gohub/pkg/helpers"

	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate app key, will print to console",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs, // 无参数
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
