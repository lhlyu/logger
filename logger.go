package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

const version = "v1.1.1"

const LOG_PREFIX = "LOGGER "
const LOG_SIGN = ">>> "

const (
	LV_DEBUG = iota
	LV_INFO
	LV_CONFIG
	LV_SIGN
	LV_ERROR
	LV_FATAL
)

const (
	SIGN_DEBUG  = "[debug ] "
	SIGN_INFO   = "[info  ] "
	SIGN_CONFIG = "[config] "
	SIGN        = "[sign  ] "
	SIGN_ERROR  = "[error ] "
	SIGN_FATAL  = "[fatal ] "
)

type Logger struct {
	mu    sync.Mutex
	Lv    int         // 日志等级
	lg    *log.Logger // 官方日志记录器
	Mod   int         // 0 - 控制台  1 - 文本
	Color int         // 0 - 颜色打印 1 - 去掉颜色
	Abs   int         // 0 - 相对路径 1 - 绝对路径
}

var _logger *Logger

var lv_color = make(map[int]int)

func init() {
	_logger = NewLogger(LV_DEBUG, "")
	fmt.Printf("logger version = %s\n", version)
	lv_color = map[int]int{
		LV_DEBUG:  Cyan,
		LV_INFO:   Blue,
		LV_CONFIG: Green,
		LV_SIGN:   Yellow,
		LV_ERROR:  Magenta,
		LV_FATAL:  Red,
	}
}

// 打印方法
func (this *Logger) print(lv int, sign string, v ...interface{}) {
	if lv < this.Lv {
		return
	}
	colorCode := lv_color[lv]
	prefix := GetColorizeFormat(ColorFormatBold, colorCode, sign)
	if this.Mod != 0 || this.Color != 0 {
		prefix = sign
	}
	this.lg.SetPrefix(prefix)
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
func printLine(s string) string {
	index := 3
	_, file, line, ok := runtime.Caller(index)
	if !ok {
		return s
	}
	if _logger.Abs == 0 {
		file = path.Base(file)
	}
	s = fmt.Sprintf("[%s:%d] %s", file, line, s)
	return s
}
