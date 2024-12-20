package make

import (
	"fmt"
	"gohub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, like: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.MinimumNArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	// 生成模型
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filepath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	// 从模板文件创建文件
	createFileFromStud(filepath, "cmd", model)

	// 友好提示
	console.Success("command name: " + model.StructName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("Please add the command to the rootCmd in main.go")
}
