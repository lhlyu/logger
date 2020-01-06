package logger

import (
	"fmt"
	_ "github.com/kataras/pio/terminal"
	"sync"
)

const (
	normal      = "%s"                  // 普通格式
	colour      = "\x1b[%dm%s\x1b[0m"   // 颜色
	highlight   = "\x1b[%d;1m%s\x1b[0m" // 高亮
	antiDisplay = "\x1b[%d;7m%s\x1b[0m" // 反显
)

type ColorVal uint32
type ColorMode uint32

const (
	ColorBlack ColorVal = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

const (
	Normal      ColorMode = iota // 普通模式
	Colour                       // 颜色模式
	Highlight                    // 高亮模式
	AntiDisplay                  // 反显模式
)

// 暴露一个
var _color = Color{
	ColorMode: Colour,
}

type Color struct {
	ColorMode ColorMode
	text      string
	value     ColorVal
	mx        sync.Mutex
}

func NewColor() *Color {
	return &Color{
		ColorMode: Colour,
	}
}

func (c Color) brush() string {
	format := ""
	switch c.ColorMode {
	case Normal:
		format = normal
		return fmt.Sprintf(format, c.text)
	case Highlight:
		format = highlight
	case AntiDisplay:
		format = antiDisplay
	default:
		format = colour
	}
	return fmt.Sprintf(format, c.value, c.text)
}

// 0 - 普通; 1 - 颜色(默认) ; 2 - 高亮 ; 3 - 反显
func (c *Color) SetColorMode(tp ColorMode) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.ColorMode = tp
}

func (c *Color) Sprint(text string) string {
	c.text = text
	return c.brush()
}

func (c *Color) Sprintf(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	return c.brush()
}

func (c *Color) Black(text string) string {
	_color.text = text
	_color.value = ColorBlack
	return _color.brush()
}

func (c *Color) Blackf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorBlack
	return _color.brush()
}

func (c *Color) Red(text string) string {
	c.text = text
	c.value = ColorRed
	return c.brush()
}

func (c *Color) Redf(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorRed
	return c.brush()
}

func (c *Color) Green(text string) string {
	c.text = text
	c.value = ColorGreen
	return c.brush()
}

func (c *Color) Greenf(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorGreen
	return c.brush()
}

func (c *Color) Yellow(text string) string {
	c.text = text
	c.value = ColorYellow
	return c.brush()
}

func (c *Color) Yellowf(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorYellow
	return c.brush()
}

func (c *Color) Blue(text string) string {
	c.text = text
	c.value = ColorBlue
	return c.brush()
}

func (c *Color) Bluef(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorBlue
	return c.brush()
}

func (c *Color) Magenta(text string) string {
	c.text = text
	c.value = ColorMagenta
	return c.brush()
}

func (c *Color) Magentaf(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorMagenta
	return c.brush()
}

func (c *Color) Cyan(text string) string {
	c.text = text
	c.value = ColorCyan
	return c.brush()
}

func (c *Color) Cyanf(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorCyan
	return c.brush()
}

func (c *Color) White(text string) string {
	_color.text = text
	_color.value = ColorWhite
	return _color.brush()
}

func (c *Color) Whitef(format string, v ...interface{}) string {
	c.text = fmt.Sprintf(format, v...)
	c.value = ColorWhite
	return c.brush()
}

/**  暴露  **/
// 0 - 普通; 1 - 颜色(默认) ; 2 - 高亮 ; 3 - 反显
func SetMode(tp ColorMode) {
	_color.mx.Lock()
	defer _color.mx.Unlock()
	_color.ColorMode = tp
}

func Sprint(text string) string {
	_color.text = text
	return _color.brush()
}

func Sprintf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	return _color.brush()
}

func Black(text string) string {
	_color.text = text
	_color.value = ColorBlack
	return _color.brush()
}

func Blackf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorBlack
	return _color.brush()
}

func Red(text string) string {
	_color.text = text
	_color.value = ColorRed
	return _color.brush()
}

func Redf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorRed
	return _color.brush()
}

func Green(text string) string {
	_color.text = text
	_color.value = ColorGreen
	return _color.brush()
}

func Greenf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorGreen
	return _color.brush()
}

func Yellow(text string) string {
	_color.text = text
	_color.value = ColorYellow
	return _color.brush()
}

func Yellowf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorYellow
	return _color.brush()
}

func Blue(text string) string {
	_color.text = text
	_color.value = ColorBlue
	return _color.brush()
}

func Bluef(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorBlue
	return _color.brush()
}

func Magenta(text string) string {
	_color.text = text
	_color.value = ColorMagenta
	return _color.brush()
}

func Magentaf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorMagenta
	return _color.brush()
}

func Cyan(text string) string {
	_color.text = text
	_color.value = ColorCyan
	return _color.brush()
}

func Cyanf(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorCyan
	return _color.brush()
}

func White(text string) string {
	_color.text = text
	_color.value = ColorWhite
	return _color.brush()
}

func Whitef(format string, v ...interface{}) string {
	_color.text = fmt.Sprintf(format, v...)
	_color.value = ColorWhite
	return _color.brush()
}
