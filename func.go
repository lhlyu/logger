package logger

import (
	"github.com/hokaccha/go-prettyjson"
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
	logger.Color = check()
	logger.pt = &prettyjson.Formatter{
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          2,
		Newline:         "\n",
	}
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

func Print(lv int, v ...interface{}) {
	printHandler(lv, "", v...)
}
func Debug(v ...interface{}) {
	printHandler(LV_DEBUG, "", v...)
}
func Info(v ...interface{}) {
	printHandler(LV_INFO, "", v...)
}
func Config(v ...interface{}) {
	printHandler(LV_CONFIG, "", v...)
}
func Sign(v ...interface{}) {
	printHandler(LV_SIGN, "", v...)
}
func Error(v ...interface{}) {
	printHandler(LV_ERROR, "", v...)
}
func Fatal(v ...interface{}) {
	printHandler(LV_FATAL, "", v...)
}
func Prompt(v ...interface{}) {
	printHandler(lv_prompt, "", v...)
}

func Printf(lv int, format string, v ...interface{}) {
	printHandler(lv, format, v...)
}
func Debugf(format string, v ...interface{}) {
	printHandler(LV_DEBUG, format, v...)
}
func Infof(format string, v ...interface{}) {
	printHandler(LV_INFO, format, v...)
}
func Configf(format string, v ...interface{}) {
	printHandler(LV_CONFIG, format, v...)
}
func Signf(format string, v ...interface{}) {
	printHandler(LV_SIGN, format, v...)
}
func Errorf(format string, v ...interface{}) {
	printHandler(LV_ERROR, format, v...)
}
func Fatalf(format string, v ...interface{}) {
	printHandler(LV_FATAL, format, v...)
}
func Promptf(format string, v ...interface{}) {
	printHandler(lv_prompt, format, v...)
}

func Json(lv int, v interface{}) {
	s, _ := _logger.pt.Marshal(v)
	printHandler(lv, "", "\n"+string(s))
}

func JsonSign(v interface{}) {
	s, _ := _logger.pt.Marshal(v)
	printHandler(LV_SIGN, "", "\n"+string(s))
}
