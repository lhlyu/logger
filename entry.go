package logger

import (
	"bytes"
	"fmt"
	"github.com/lhlyu/logger/color"
	"time"
)

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logx(FatalLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.logx(FatalLevel, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logx(ErrorLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorln(v ...interface{}) {
	l.logx(ErrorLevel, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logx(WarnLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Warnln(v ...interface{}) {
	l.logx(WarnLevel, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logx(InfoLevl, fmt.Sprintf(format, v...))
}

func (l *Logger) Infoln(v ...interface{}) {
	l.logx(InfoLevl, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logx(DebugLevl, fmt.Sprintf(format, v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.logx(DebugLevl, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logx(DisableLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.logx(DisableLevel, v...)
}

func (l *Logger) logx(level Level, v ...interface{}) {
	if l.Level < level || len(v) == 0 {
		return
	}
	ctx := newCtx(l, v)
	if l.Before != nil {
		l.Before(ctx)
	}
	if l.After != nil {
		defer l.After(ctx)
	}

	buf := bytes.Buffer{}
	if l.Color.ColorMode != color.Normal {
		if isTerminal(l.Out) {
			meta := parseMeta(l.Color, level)
			buf.WriteString(meta.cl(meta.LvName))
		}
	}
	if l.TimeFormat != "" {
		buf.WriteString(time.Now().Format(l.TimeFormat))
	}
	var (
		byts []byte
		err  error
	)
	if len(v) > 1 {
		byts, err = l.Formatter.Format(v)
	} else {
		byts, err = l.Formatter.Format(v[0])
	}
	if err != nil {
		ctx.Err = err
		return
	}
	if l.Formatter.NewLine() {
		buf.WriteString("\n")
	}
	buf.Write(byts)
	l.lg.Output(2, buf.String())
}
