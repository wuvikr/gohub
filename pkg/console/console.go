package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Success(message string) {
	colorOut(message, "green")
}

func Warning(message string) {
	colorOut(message, "yellow")
}

func Error(message string) {
	colorOut(message, "red")
}

func Exit(message string) {
	colorOut(message, "red")
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func colorOut(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}
