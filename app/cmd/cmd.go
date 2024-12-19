package cmd

import (
	"gohub/pkg/helpers"
	"os"

	"github.com/spf13/cobra"
)

var Env string

// RegisterGlobalFlags 注册全局标志
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, e.g. --env=testing will load .env.testing file")
}

func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])

	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
