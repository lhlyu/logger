package logger

import (
	isatty "github.com/mattn/go-isatty"
	"os"
)

func check() int {
	if os.Getenv("OS") == "Windows_NT" && isatty.IsTerminal(os.Stdout.Fd()) {
		return 1
	}
	return 0
}
