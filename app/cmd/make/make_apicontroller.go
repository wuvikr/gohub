package make

import (
	"fmt"
	"gohub/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create a new API controller, e.g. make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	// 参数处理，要求附带 API 版本号
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	// apiVersion 用来拼接目标路径
	// name 用来生成 cmd.Model 实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标路径
	filepath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	// 生成文件
	createFileFromStub(filepath, "apicontroller", model)
}
