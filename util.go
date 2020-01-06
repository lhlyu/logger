package logger

import (
	"io"
	"os"
)

func isTerminal(f io.Writer) bool {
	switch v := f.(type) {
	case *os.File:
		if v.Name() == "/dev/stdout" || v.Name() == "/dev/stderr" {
			return true
		}
	}
	return false
}
