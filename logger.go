package logger

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"log"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"sync"
)

const version = "v1.2.1"

const LOG_PREFIX = "LOGGER "
const LOG_SIGN = ">>> "

const (
	LV_DEBUG = iota
	LV_INFO
	LV_CONFIG
	LV_SIGN
	LV_ERROR
	LV_FATAL
	lv_prompt
)

const (
	SIGN_DEBUG  = "[debug ] "
	SIGN_INFO   = "[info  ] "
	SIGN_CONFIG = "[config] "
	SIGN        = "[sign  ] "
	SIGN_ERROR  = "[error ] "
	SIGN_FATAL  = "[fatal ] "
	sign_prompt = ""
)

var lvSignMap = map[int]string{
	LV_DEBUG:  SIGN_DEBUG,
	LV_INFO:   SIGN_INFO,
	LV_CONFIG: SIGN_CONFIG,
	LV_SIGN:   SIGN,
	LV_ERROR:  SIGN_ERROR,
	LV_FATAL:  SIGN_FATAL,
	lv_prompt:  sign_prompt,
}

type Logger struct {
	mu    sync.Mutex
	Lv    int                   // 日志等级
	lg    *log.Logger           // 官方日志记录器
	Mod   int                   // 0 - 控制台   1 - 文本
	Color int                   // 0 - 颜色打印 1 - 去掉颜色
	Abs   int                   // 0 - 相对路径 1 - 绝对路径
	pt    *prettyjson.Formatter // 格式化工具
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
		lv_prompt: Black,
	}
}

// 打印方法
func (this *Logger) print(lv int, sign string, s string) {
	if lv < this.Lv {
		return
	}
	colorCode := lv_color[lv]
	prefix := GetColorizeFormat(ColorFormatBold, colorCode, sign)
	if this.Mod != 0 || this.Color != 0 {
		prefix = sign
	}
	this.lg.SetPrefix(prefix)
	this.lg.Println(s)
	if lv == LV_FATAL {
		os.Exit(1)
	}
}

func printHandler(lv int, format string, v ...interface{}) {
	var line string // 方法和行号
	var lineFormat = " %s %s"
	var params []interface{}
	if lv >= LV_SIGN && lv < lv_prompt{
		lineFormat = "[ %s ] %s %s"
		stacks := string(debug.Stack())
		stackArr := strings.Split(stacks, "\n")
		line = strings.TrimSpace(stackArr[8])
		lineArr := strings.Split(line, " ")
		line = lineArr[0]
		if _logger.Abs == 0 {
			line = path.Base(line)
		}
		params = append(params, line)
	}
	var s string
	if format == "" {
		s = fmt.Sprint(v...)
	} else {
		s = fmt.Sprintf(format, v...)
	}
	params = append(params, LOG_SIGN)
	params = append(params, s)
	s = fmt.Sprintf(lineFormat, params...)
	sign, ok := lvSignMap[lv]
	if !ok {
		sign = SIGN_DEBUG
	}
	_logger.print(lv, sign, s)
}
