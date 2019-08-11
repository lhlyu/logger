package logger

import (
	"fmt"
)

const (
	ColorFormat          = "\x1b[%dm%s\x1b[0m"   // 普通格式
	ColorFormatBold      = "\x1b[%d;1m%s\x1b[0m" // 加粗格式
	ColorFormatUnderline = "\x1b[%d;4m%s\x1b[0m" // 下划线格式
)

// 字体色
const (
	ColorBlack   = 30
	ColorRed     = 31
	ColorGreen   = 32
	ColorYellow  = 33
	ColorBlue    = 34
	ColorMagenta = 35 // 品红
	ColorCyan    = 36 // 青色
	ColorWhite   = 37
)

// 背景色
const (
	BgColorBlack   = 40
	BgColorRed     = 41
	BgColorGreen   = 42
	BgColorYellow  = 43
	BgColorBlue    = 44
	BgColorMagenta = 45 // 品红
	BgColorCyan    = 46 // 青色
	BgColorWhite   = 47
)

//  ANSI控制码:
//
//  QUOTE:
//  \x1b[0m     关闭所有属性
//  \x1b[1m     设置高亮度
//  \x1b[4m     下划线
//  \x1b[5m     闪烁
//  \x1b[7m     反显
//  \x1b[8m     消隐
//  \x1b[30m   --  \x1b[37m   设置前景色
//  \x1b[40m   --  \x1b[47m   设置背景色
//  \x1b[nA    光标上移n行
//  \x1b[nB    光标下移n行
//  \x1b[nC    光标右移n行
//  \x1b[nD    光标左移n行
//  \x1b[y;xH  设置光标位置
//  \x1b[2J    清屏
//  \x1b[K     清除从光标到行尾的内容
//  \x1b[s     保存光标位置
//  \x1b[u     恢复光标位置
//  \x1b[?25l  隐藏光标
//  \x1b[?25h  显示光标

func colorize(colorFormat string, colorCode int, s string) string {
	return fmt.Sprintf(colorFormat, colorCode, s)
}

func colorFormat(colorCode int, s string) string {
	return colorize(ColorFormat, colorCode, s)
}

func colorFormatBold(colorCode int, s string) string {
	return colorize(ColorFormatBold, colorCode, s)
}

func colorFormatUnderline(colorCode int, s string) string {
	return colorize(ColorFormatUnderline, colorCode, s)
}

// @方法

// 普通红色字体
func Red(s string) string {
	return colorFormat(ColorRed, s)
}

// 普通绿色字体
func Green(s string) string {
	return colorFormat(ColorGreen, s)
}

// 普通黄色字体
func Yellow(s string) string {
	return colorFormat(ColorYellow, s)
}

// 普通蓝色字体
func Blue(s string) string {
	return colorFormat(ColorBlue, s)
}

// 普通品红字体
func Magenta(s string) string {
	return colorFormat(ColorMagenta, s)
}

// 普通青色字体
func Cyan(s string) string {
	return colorFormat(ColorCyan, s)
}

// 红色下划线字体
func RedLine(s string) string {
	return colorFormatUnderline(ColorRed, s)
}

// 绿色下划线字体
func GreenLine(s string) string {
	return colorFormatUnderline(ColorGreen, s)
}

// 黄色下划线字体
func YellowLine(s string) string {
	return colorFormatUnderline(ColorYellow, s)
}

// 蓝色下划线字体
func BlueLine(s string) string {
	return colorFormatUnderline(ColorBlue, s)
}

// 品红下划线字体
func MagentaLine(s string) string {
	return colorFormatUnderline(ColorMagenta, s)
}

// 青色下划线字体
func CyanLine(s string) string {
	return colorFormatUnderline(ColorCyan, s)
}

// 粗体红色字体
func RedBlod(s string) string {
	return colorFormatBold(ColorRed, s)
}

// 粗体绿色字体
func GreenBlod(s string) string {
	return colorFormatBold(ColorGreen, s)
}

// 粗体黄色字体
func YellowBlod(s string) string {
	return colorFormatBold(ColorYellow, s)
}

// 粗体蓝色字体
func BlueBlod(s string) string {
	return colorFormatBold(ColorBlue, s)
}

// 粗体品红字体
func MagentaBlod(s string) string {
	return colorFormatBold(ColorMagenta, s)
}

// 粗体青色字体
func CyanBlod(s string) string {
	return colorFormatBold(ColorCyan, s)
}

// 普通红色背景字体
func RedBg(s string) string {
	return colorFormat(BgColorRed, s)
}

// 普通绿色背景字体
func GreenBg(s string) string {
	return colorFormat(BgColorGreen, s)
}

// 普通黄色背景字体
func YellowBg(s string) string {
	return colorFormat(BgColorYellow, s)
}

// 普通蓝色背景字体
func BlueBg(s string) string {
	return colorFormat(BgColorBlue, s)
}

// 普通品红背景字体
func MagentaBg(s string) string {
	return colorFormat(BgColorMagenta, s)
}

// 普通青色背景字体
func CyanBg(s string) string {
	return colorFormat(BgColorCyan, s)
}

// 红色下划线背景字体
func RedLineBg(s string) string {
	return colorFormatUnderline(BgColorRed, s)
}

// 绿色下划线背景字体
func GreenLineBg(s string) string {
	return colorFormatUnderline(BgColorGreen, s)
}

// 黄色下划线背景字体
func YellowLineBg(s string) string {
	return colorFormatUnderline(BgColorYellow, s)
}

// 蓝色下划线背景字体
func BlueLineBg(s string) string {
	return colorFormatUnderline(BgColorBlue, s)
}

// 品红下划线背景字体
func MagentaLineBg(s string) string {
	return colorFormatUnderline(BgColorMagenta, s)
}

// 青色下划线背景字体
func CyanLineBg(s string) string {
	return colorFormatUnderline(BgColorCyan, s)
}

// 粗体红色背景字体
func RedBlodBg(s string) string {
	return colorFormatBold(BgColorRed, s)
}

// 粗体绿色背景字体
func GreenBlodBg(s string) string {
	return colorFormatBold(BgColorGreen, s)
}

// 粗体黄色背景字体
func YellowBlodBg(s string) string {
	return colorFormatBold(BgColorYellow, s)
}

// 粗体蓝色背景字体
func BlueBlodBg(s string) string {
	return colorFormatBold(BgColorBlue, s)
}

// 粗体品红背景字体
func MagentaBlodBg(s string) string {
	return colorFormatBold(BgColorMagenta, s)
}

// 粗体青色背景字体
func CyanBlodBg(s string) string {
	return colorFormatBold(BgColorCyan, s)
}
