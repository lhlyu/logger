package logger

import "fmt"

/**
30-37 设置文本颜色
black: 30
red: 31
green: 32
yellow: 33
blue: 34
magenta: 35
cyan: 36
white: 37
40–47 设置文本背景颜色
39 重置文本颜色
49 重置背景颜色
1 加粗文本 / 高亮
22 重置加粗 / 高亮
0 重置所有文本属性（颜色，背景，亮度等）为默认值
*/

// 输出格式
const (
	ColorFormat      = "\x1b[%dm%s\x1b[0m"   // 普通格式
	ColorFormatBold  = "\x1b[%d;1m%s\x1b[0m" // 加粗格式
)

// 字体色
const (
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35 // 品红
	Cyan    = 36 // 青色
	White   = 37
)

// 背景色
const (
	BlackBackground   = 40
	RedBackground     = 41
	GreenBackground   = 42
	YellowBackground  = 43
	BlueBackground    = 44
	MagentaBackground = 45 // 品红
	CyanBackground    = 46 // 青色
	WhiteBackground   = 47
)

func GetColorize(colorCode int, s string) string {
	return GetColorizeFormat(ColorFormat, colorCode, s)
}

func GetColorizeFormat(colorFormat string, colorCode int, s string) string {
	return fmt.Sprintf(colorFormat, colorCode, s)
}
