package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Create a model, should be snake_case, like: make model user",
	Run:   runMakeModel,
	Args:  cobra.MinimumNArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	// 生成模型
	model := makeModelFromString(args[0])

	// 确保目标文件路径存在
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	os.MkdirAll(dir, os.ModePerm)

	// 替换变量
	createFileFromStud(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStud(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStud(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
}
