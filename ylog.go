package logger

import (
	"fmt"
	"github.com/lhlyu/logger/pio"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type yLog struct {
	mu         sync.Mutex
	Printer    *pio.Printer
	TimeFormat string    // 时间格式
	Level      int       // 日志等级
	ColorMod   bool      // 颜色模式  终端默认开启
	Delims     [2]string // 标签左右符号，默认是 [ 和 ]
	Location   int       // 打印调用栈信息 默认关闭  0 - 不打印  1- 打印调用方法名和行号  2-打印文件路径和行号
}

var _ylog *yLog

// 初始化
func init() {
	_ylog = new()
	fmt.Println(CyanBlod("logger's version = " + version))
}

// 新建
func new() *yLog {
	return &yLog{
		mu:         sync.Mutex{},
		Printer:    pio.NewPrinter("default", os.Stdout),
		TimeFormat: "2006-01-02 15:04:05",
		Level:      1,
		ColorMod:   true,
		Delims:     [2]string{"[", "]"},
		Location:   0,
	}
}

// 打印文本
func print(level int, s string) {
	if level < _ylog.Level {
		return
	}
	md := getLevelMetadata(level)
	_ylog.Printer.Marshal(pio.Text)
	pf := prefix(md)
	if _ylog.Location != 0 {
		fp, fn, line := getCaller()
		switch _ylog.Location {
		case 1:
			pf += fmt.Sprintf("[%s:%d] ", fn, line)
		case 2:
			pf += fmt.Sprintf("[%s:%d] ", fp, line)
		}
	}
	s = pf + s
	_ylog.Printer.Println(s)
	if md.NextHandler != nil{
		md.NextHandler()
	}
}

func printJson(level int, v interface{}) {
	if level < _ylog.Level {
		return
	}
	md := getLevelMetadata(level)
	_ylog.Printer.Marshal(pio.JSONIndent)
	pf := prefix(md)
	if _ylog.Location != 0 {
		fp, fn, line := getCaller()
		switch _ylog.Location {
		case 1:
			pf += fmt.Sprintf("[%s:%d] ", fn, line)
		case 2:
			pf += fmt.Sprintf("[%s:%d] ", fp, line)
		}
	}
	_ylog.Printer.Println(pf)
	_ylog.Printer.Println(v)
	if md.NextHandler != nil{
		md.NextHandler()
	}
}

func printNormal(v ...interface{}) {
	_ylog.Printer.Marshal(pio.JSON)
	s := fmt.Sprint(v...)
	if _ylog.ColorMod && _ylog.Printer.IsTerminal {
		s = CyanBlod(s)
	}
	_ylog.Printer.Println(s)
}

func getCaller() (string, string, int) {
	var pcs [10]uintptr
	n := runtime.Callers(2, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "github.com/lhlyu/logger") {
			return frame.File, frame.Func.Name(), frame.Line
		}
		if !more {
			break
		}
	}
	return "?", "?", 0
}

func prefix(md levelMetadata) string {
	now := time.Now().Format("[" + _ylog.TimeFormat + "] ")
	if md.Level == LevelNone {
		return now
	}
	txt := _ylog.Delims[0] + md.Txt + _ylog.Delims[1]
	if _ylog.ColorMod && _ylog.Printer.IsTerminal {
		if md.ColorHanler != nil {
			txt = md.ColorHanler(txt)
		}
	}
	return txt + " " + now
}
