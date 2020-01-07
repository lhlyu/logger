package logger

import "github.com/lhlyu/logger/color"

type levelMeta struct {
	LvName string
	cl     func(text string) string
}

func parseMeta(c *color.Color, lv Level) *levelMeta {
	switch lv {
	case FatalLevel:
		return &levelMeta{
			LvName: "[fatal] ",
			cl:     c.Magenta,
		}
	case ErrorLevel:
		return &levelMeta{
			LvName: "[error] ",
			cl:     c.Red,
		}
	case WarnLevel:
		return &levelMeta{
			LvName: "[warn ] ",
			cl:     c.Yellow,
		}
	case InfoLevl:
		return &levelMeta{
			LvName: "[info ] ",
			cl:     c.Blue,
		}
	case DebugLevl:
		return &levelMeta{
			LvName: "[debug] ",
			cl:     c.Green,
		}
	}
	return &levelMeta{
		LvName: "",
		cl: func(text string) string {
			return text
		},
	}
}
