package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create a new factory, e.g. make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1),
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	// 格式化模型
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filepath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	// 生成文件
	createFileFromStub(filepath, "factory", model)
}
