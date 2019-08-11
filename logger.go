package logger

import (
	"fmt"
	"io"
	"os"
)

const version = "v2.0.0"

// 重置,回归初始
func Reset() {
	_ylog = new()
}

// 设置输出端
func SetWriters(writers ...io.Writer) {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	_ylog.Printer.SetOutput()
}

// 设置时间格式
func SetTimeFormat(s string) {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	_ylog.TimeFormat = s
}

// 设置文件路径
func SetFilePath(s string) error {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	f, err := os.OpenFile(s, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_ylog.Printer.SetOutput(f)
	return nil
}

// 设置日志等级
func SetLevel(level int) {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	_ylog.Level = level
}

// 设置是否启用颜色打印
func SetColorMod(enable bool) {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	_ylog.ColorMod = enable
}

// 设置日志左右标签 默认是 [  和  ]
func SetDelims(left, right string) {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	_ylog.Delims = [2]string{left, right}
}

// 设置日志打印文件和行号  0 - 不打印(默认)  1- 打印调用方法名和行号  2-打印文件路径和行号
func SetLocation(location int) {
	_ylog.mu.Lock()
	defer _ylog.mu.Unlock()
	_ylog.Location = location
}

// 无前缀打印信息
func PNone(v ...interface{}) {
	print(LevelNone, fmt.Sprint(v...))
}

// 打印致命信息，会退出程序
func PFatal(v ...interface{}) {
	print(LevelFatal, fmt.Sprint(v...))
}

// 打印错误信息
func PError(v ...interface{}) {
	print(LevelError, fmt.Sprint(v...))
}

// 打印警告信息
func PWarn(v ...interface{}) {
	print(LevelWarn, fmt.Sprint(v...))
}

// 打印普通信息
func PInfo(v ...interface{}) {
	print(LevelInfo, fmt.Sprint(v...))
}

// 打印调试信息
func PDebug(v ...interface{}) {
	print(LevelDebug, fmt.Sprint(v...))
}

// 格式化打印无前缀信息
func PNonef(format string, v ...interface{}) {
	print(LevelNone, fmt.Sprintf(format, v...))
}

// 格式化打印致命信息，会退出程序
func PFatalf(format string, v ...interface{}) {
	print(LevelFatal, fmt.Sprintf(format, v...))
}

// 格式化打印错误信息
func PErrorf(format string, v ...interface{}) {
	print(LevelError, fmt.Sprintf(format, v...))
}

// 格式化打印警告信息
func PWarnf(format string, v ...interface{}) {
	print(LevelWarn, fmt.Sprintf(format, v...))
}

// 格式化打印普通信息
func PInfof(format string, v ...interface{}) {
	print(LevelInfo, fmt.Sprintf(format, v...))
}

// 格式化打印调试信息
func PDebugf(format string, v ...interface{}) {
	print(LevelDebug, fmt.Sprintf(format, v...))
}

// 打印一个对象
func PJson(v interface{}) {
	printJson(LevelInfo, v)
}

// 自定义等级，打印一个对象
func PJsonL(level int, v interface{}) {
	printJson(level, v)
}

// 一个普通打印
func PNormal(v ...interface{}) {
	printNormal(v...)
}

// 一个普通格式化打印
func PNormalf(format string, v ...interface{}) {
	printNormal(fmt.Sprintf(format, v...))
}
