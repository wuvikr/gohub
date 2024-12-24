package make

import (
	"embed"
	"gohub/pkg/console"
	"gohub/pkg/file"
	"gohub/pkg/str"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

type Model struct {
	TbaleName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// Add subcommands
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeModel,
		CmdMakeAPIController,
		CmdMakeRequest,
	)
}

// makeModelFromString 格式化用户输入的字符串
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TbaleName = str.Snake(str.Plural(name))
	model.VariableName = str.LowerCamel(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	model.PackageName = str.Snake(model.StructName)
	return model
}

// createFileFromStud 从模板文件创建文件
func createFileFromStud(filePath string, stubName string, model Model, variables ...interface{}) {

	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// 目标文件已存在
	if file.Exist(filePath) {
		console.Exit(filePath + " already exists")
	}

	// 读取模板文件
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	// 添加默认的替换变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{TbaleName}}"] = model.TbaleName
	replaces["{{PackageName}}"] = model.PackageName

	// 对模板进行替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	// 写入文件
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	console.Success(filePath + " created successfully")
}
