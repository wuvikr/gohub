package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create a new request, e.g. make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	name := args[0]
	model := makeModelFromString(name)

	// 组建目标路径
	filepath := fmt.Sprintf("app/http/requests/%s_request.go", model.TbaleName)

	// 生成文件
	createFileFromStub(filepath, "request", model)
}
