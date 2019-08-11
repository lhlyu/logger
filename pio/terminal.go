package pio

import (
	"github.com/lhlyu/logger/pio/terminal"
	"io"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func isTerminal(output io.Writer) bool {
	isTerminal := !IsNop(output) || terminal.IsTerminal(output)

	if !isTerminal || runtime.GOOS != "windows" {
		return isTerminal
	}

	cmd := exec.Command("cmd", "ver")

	b, err := cmd.Output()
	if err != nil {
		return false
	}

	lines := string(b)
	if lines == "" {
		return false
	}

	start := strings.IndexByte(lines, '[')
	end := strings.IndexByte(lines, ']')

	winLine := lines[start+1 : end]
	if len(winLine) < 10 {
		return false
	}

	versionsLine := winLine[strings.IndexByte(winLine, ' ')+1:]

	versionSems := strings.Split(versionsLine, ".")

	if len(versionSems) < 3 {
		return false
	}

	if versionSems[0] != "10" {
		return false
	}

	buildNumber, err := strconv.Atoi(versionSems[2])

	if err != nil {
		return false
	}

	return buildNumber >= 10586
}
