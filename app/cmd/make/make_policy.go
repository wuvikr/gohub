package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create a new policy, e.g. make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	// 格式化模型
	model := makeModelFromString(args[0])

	// 确保目标文件路径存在
	os.MkdirAll("app/policies", os.ModePerm)
	// 拼接目标文件路径
	filepath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	// 从模板文件创建文件
	createFileFromStub(filepath, "policy", model)
}
