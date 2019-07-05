package logger

import (
	"os"
	isatty "github.com/mattn/go-isatty"
)

func check() int{
	if os.Getenv("OS") == "Windows_NT" && isatty.IsTerminal(os.Stdout.Fd()){
		return 1
	}
	return 0
}
