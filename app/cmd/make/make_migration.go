package make

import (
	"fmt"
	"gohub/pkg/app"
	"gohub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a new migration file, e.g. make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1),
}

func runMakeMigration(cmd *cobra.Command, args []string) {

	// 格式化时间
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStud(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created, after modify it, use 'gohub migrate up' to run it")
}
