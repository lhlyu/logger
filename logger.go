package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const version = "v1.0.7"

const LOG_PREFIX = "LOGGER "
const LOG_SIGN = ">>> "

const (
	LV_SIGN  = -1
	LV_DEBUG = iota - 1
	LV_FUNC
	LV_INFO
	LV_CONFIG
	LV_ERROR
	LV_FATAL
)

const (
	SIGN        = "[sign  ] "
	SIGN_DEBUG  = "[debug ] "
	SIGN_FUNC   = "[func  ] "
	SIGN_INFO   = "[info  ] "
	SIGN_CONFIG = "[config] "
	SIGN_ERROR  = "[error ] "
	SIGN_FATAL  = "[fatal ] "
)

var lv_color = map[int]string{
	LV_SIGN:   "%s",
	LV_DEBUG:  "%s",
	LV_FUNC:   "\x1B[33m%s\x1b[0m",
	LV_INFO:   "\x1B[34m%s\x1b[0m",
	LV_CONFIG: "\x1b[32m%s\x1b[0m",
	LV_ERROR:  "\x1b[95m%s\x1b[0m",
	LV_FATAL:  "\x1b[91m%s\x1b[0m",
}

type Logger struct {
	Lv  int         // 日志等级
	lg  *log.Logger // 官方日志记录器
	mod int         // 0 - 控制台  1 - 文本
}

var _logger *Logger
var _color = 0
var _abs = 0

func init() {
	_logger = NewLogger(LV_DEBUG, "")
	fmt.Printf("logger version = %s\n", version)
}

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
	logger.mod = mod
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
	_color = open
}
// 设置等级
func SetLevel(level int) {
	_logger.Lv = level
}
// 设置文件定位是否开启绝对路径,0 = 不开启  其他 = 开启
func SetAbs(abs int) {
	_abs = abs
}

// 打印方法
func (this *Logger) print(lv int, sign string, v ...interface{}) {
	if lv < this.Lv {
		return
	}
	format := lv_color[lv]
	if this.mod != 0 || _color != 0 {
		format = "%s"
	}
	this.lg.SetPrefix(fmt.Sprintf(format, sign))
	s := LOG_SIGN + fmt.Sprint(v...)
	if lv == LV_SIGN {
		s = printLine(s)
	}
	if lv >= LV_ERROR {
		s = printLine(s)
	}
	this.lg.Println(s)
	if lv == LV_FATAL {
		os.Exit(1)
	}
}
// 打印文件和所在行
func printLine(s string) string{
	index := 2
	for {
		_, file, line, ok := runtime.Caller(index)
		if !ok {
			break
		}
		if _abs == 0{
			file = path.Base(file)
		}
		if strings.LastIndex(file, "logger.go") == -1 {
			s = fmt.Sprintf("[%s:%d] %s", file, line, s)
			break
		}
		index += 1
	}
	return s
}

func Sign(v ...interface{}){
	_logger.print(LV_SIGN, SIGN,v...)
}
func Debug(v ...interface{}) {
	_logger.print(LV_DEBUG, SIGN_DEBUG, v...)
}
func Func(v ...interface{}) {
	_logger.print(LV_FUNC, SIGN_FUNC, v...)
}
func Info(v ...interface{}) {
	_logger.print(LV_INFO, SIGN_INFO, v...)
}
func Config(v ...interface{}) {
	_logger.print(LV_CONFIG, SIGN_CONFIG, v...)
}
func Error(v ...interface{}) {
	_logger.print(LV_ERROR, SIGN_ERROR, v...)
}
func Fatal(v ...interface{}) {
	_logger.print(LV_FATAL, SIGN_FATAL, v...)
}

func Signf(format string, v ...interface{}) {
	Sign(fmt.Sprintf(format, v...))
}
func Debugf(format string, v ...interface{}) {
	Debug(fmt.Sprintf(format, v...))
}
func Funcf(format string, v ...interface{}) {
	Func(fmt.Sprintf(format, v...))
}
func Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v...))
}
func Configf(format string, v ...interface{}) {
	Config(fmt.Sprintf(format, v...))
}
func Errorf(format string, v ...interface{}) {
	Error(fmt.Sprintf(format, v...))
}
func Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v...))
}
