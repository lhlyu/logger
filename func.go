package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func NewLogger(lv int, fldir string) *Logger {
	var lg *log.Logger
	var mod int
	if fldir != "" {
		now := time.Now()
		flname := now.Format("20060102.log")
		fl, err := os.Create(path.Join(fldir, flname))
		if err != nil {
			Fatal(err)
		}
		lg = log.New(fl, LOG_PREFIX, 3)
		mod = 1
	} else {
		lg = log.New(os.Stdout, LOG_PREFIX, 3)
	}
	logger := &Logger{}
	logger.Lv = lv
	logger.Mod = mod
	logger.lg = lg
	return logger
}

// 设置日志记录器
func SetLogger(logger *Logger) {
	if logger != nil {
		_logger = logger
	}
}

// 设置是否开启颜色, 0 = 开启   其他 = 不开启
func SetColor(open int) {
	_logger.mu.Lock()
	_logger.Color = open
	_logger.mu.Unlock()
}

// 设置等级
func SetLevel(level int) {
	_logger.mu.Lock()
	_logger.Lv = level
	_logger.mu.Unlock()
}

// 设置文件定位是否开启绝对路径,0 = 不开启  其他 = 开启
func SetAbs(abs int) {
	_logger.mu.Lock()
	_logger.Abs = abs
	_logger.mu.Unlock()
}

func Debug(v ...interface{}) {
	_logger.print(LV_DEBUG, SIGN_DEBUG, v...)
}
func Info(v ...interface{}) {
	_logger.print(LV_INFO, SIGN_INFO, v...)
}
func Config(v ...interface{}) {
	_logger.print(LV_CONFIG, SIGN_CONFIG, v...)
}
func Sign(v ...interface{}) {
	_logger.print(LV_SIGN, SIGN, v...)
}
func Error(v ...interface{}) {
	_logger.print(LV_ERROR, SIGN_ERROR, v...)
}
func Fatal(v ...interface{}) {
	_logger.print(LV_FATAL, SIGN_FATAL, v...)
}

func Debugf(format string, v ...interface{}) {
	Debug(fmt.Sprintf(format, v...))
}
func Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v...))
}
func Configf(format string, v ...interface{}) {
	Config(fmt.Sprintf(format, v...))
}
func Signf(format string, v ...interface{}) {
	Sign(fmt.Sprintf(format, v...))
}
func Errorf(format string, v ...interface{}) {
	Error(fmt.Sprintf(format, v...))
}
func Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v...))
}
