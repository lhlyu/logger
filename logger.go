package logger

import (
	"context"
	"github.com/lhlyu/logger/color"
	"io"
	"log"
	"os"
	"sync"
)

const version = "v3.0.0"

const (
	default_time_format = "2006-01-02 15:04:05 ▶ "
)

type Before func(ctx *Ctx)
type After func(ctx *Ctx)

type Logger struct {
	Out        io.Writer    // 输出流
	Before     Before       // 前置处理器
	After      After        // 后置处理器
	Formatter  Formatter    // 内容格式化
	TimeFormat string       // 时间格式化
	Level      Level        // 等级
	Color      *color.Color // 颜色控制
	Context    context.Context
	mx         sync.Mutex
	lg         *log.Logger
}

func New() *Logger {
	return &Logger{
		Out:        os.Stdout,
		Formatter:  new(textFormatter),
		Level:      InfoLevl,
		TimeFormat: default_time_format,
		Color:      color.NewColor(),
		lg:         log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) SetOutput(output io.Writer) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.Out = output
	return l
}

func (l *Logger) SetLevel(lv string) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.Level = ParseLevel(lv)
	return l
}

func (l *Logger) SetFormatter(formatter Formatter) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.Formatter = formatter
	return l
}

func (l *Logger) SetTimeFormat(timeFormat string) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.TimeFormat = timeFormat
	return l
}

func (l *Logger) WithContext(context context.Context) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.Context = context
	return l
}

func (l *Logger) SetColorMode(colorMode color.ColorMode) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	if l.Color == nil {
		l.Color = &color.Color{}
	}
	l.Color.ColorMode = colorMode
	return l
}

func (l *Logger) AddBefore(befor Before) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	if l.Before == nil {
		l.Before = befor
		return l
	}
	preBefore := l.Before
	nextBefore := befor
	l.Before = func(ctx *Ctx) {
		preBefore(ctx)
		if !ctx.stop {
			nextBefore(ctx)
		}
	}
	return l
}

func (l *Logger) AddAfter(after After) *Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	if l.After == nil {
		l.After = after
		return l
	}
	preAfter := l.After
	nextAfter := after
	l.After = func(ctx *Ctx) {
		preAfter(ctx)
		if !ctx.stop {
			nextAfter(ctx)
		}
	}
	return l
}
