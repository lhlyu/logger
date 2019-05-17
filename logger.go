package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

const version = "v1.0.3"

const LOG_PREFIX = "LOGGER "
const LOG_SIGN = ">>> "

const (
	LV_DEBUG = iota
	LV_FUNC
	LV_INFO
	LV_CONFIG
	LV_ERROR
	LV_FATAL
)

const (
	SIGN_DEBUG  = "[debug   ] "
	SIGN_FUNC   = "[func    ] "
	SIGN_INFO   = "[info    ] "
	SIGN_CONFIG = "[config  ] "
	SIGN_ERROR  = "[error   ] "
	SIGN_FATAL  = "[fatal   ] "
)

var lv_color = map[int]string{
	LV_DEBUG:"%s",
	LV_FUNC:"\x1B[33m%s\x1b[0m",
	LV_INFO:"\x1B[34m%s\x1b[0m",
	LV_CONFIG:"\x1b[32m%s\x1b[0m",
	LV_ERROR:"\x1b[95m%s\x1b[0m",
	LV_FATAL:"\x1b[91m%s\x1b[0m",
}

type Logger struct {
	Lv  int                // 日志等级
	lg  *log.Logger        // 官方日志记录器
	mod int               // 0 - 控制台  1 - 文本
}

var _logger *Logger
var _color = 0

func init(){
	_logger = NewLogger(LV_DEBUG,"")
}

func NewLogger(lv int,fldir string) *Logger{
	var lg *log.Logger
	var mod int
	if fldir != ""{
		now := time.Now()
		flname := now.Format("20060102.log")
		fl,err := os.Create(path.Join(fldir,flname))
		if err != nil{
			Fatal(err)
		}
		lg = log.New(fl,LOG_PREFIX,3)
		mod = 1
	}else{
		lg = log.New(os.Stdout,LOG_PREFIX,3)
	}
	logger := &Logger{}
	logger.Lv = lv
	logger.mod = mod
	logger.lg = lg
	return logger
}

func SetLogger(logger *Logger){
	if logger != nil {
		_logger = logger
	}
}

func SetColor(open int){
	_color = open

}

func (this *Logger) print(lv int,sign string,v ...interface{}){
	if lv < this.Lv{
		return
	}
	format := lv_color[lv]
	if this.mod != 0 || _color != 0{
		format = "%s"
	}
	this.lg.SetPrefix(fmt.Sprintf(format,sign))
	s := LOG_SIGN + fmt.Sprint(v...)
	this.lg.Println(s)
	if lv == LV_FATAL{
		os.Exit(1)
	}
}

func Debug(v ...interface{}){
	_logger.print(LV_DEBUG,SIGN_DEBUG,v...)
}
func Func(v ...interface{}){
	_logger.print(LV_FUNC,SIGN_FUNC,v...)
}
func Info(v ...interface{}){
	_logger.print(LV_INFO,SIGN_INFO,v...)
}
func Config(v ...interface{}){
	_logger.print(LV_CONFIG,SIGN_CONFIG,v...)
}
func Error(v ...interface{}){
	_logger.print(LV_ERROR,SIGN_ERROR,v...)
}
func Fatal(v ...interface{}){
	_logger.print(LV_FATAL,SIGN_FATAL,v...)
}

func Debugf(format string,v ...interface{}){
	Debug(fmt.Sprintf(format,v...))
}
func Funcf(format string,v ...interface{}){
	Func(fmt.Sprintf(format,v...))
}
func Infof(format string,v ...interface{}){
	Info(fmt.Sprintf(format,v...))
}
func Configf(format string,v ...interface{}){
	Config(fmt.Sprintf(format,v...))
}
func Errorf(format string,v ...interface{}){
	Error(fmt.Sprintf(format,v...))
}
func Fatalf(format string,v ...interface{}){
	Fatal(fmt.Sprintf(format,v...))
}