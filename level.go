package logger

import "strings"

type Level uint32

const (
	FatalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevl
	DebugLevl

	DisableLevel
)

func ParseLevel(lv string) Level {
	switch strings.ToLower(lv) {
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevl
	case "debug":
		return DebugLevl
	default:
		return DisableLevel
	}
}
