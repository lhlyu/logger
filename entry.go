package logger

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

type Entry struct {
	logger *Logger
	lg     *log.Logger
}

func NewEntry(logger *Logger) *Entry {
	lg := log.New(logger.Out, "", 0)
	return &Entry{
		logger: logger,
		lg:     lg,
	}
}

func (entry *Entry) Fatalf(format string, v ...interface{}) {
	entry.logx(FatalLevel, fmt.Sprintf(format, v...))
}

func (entry *Entry) Fatalln(v ...interface{}) {
	entry.logx(FatalLevel, v...)
}

func (entry *Entry) Errorf(format string, v ...interface{}) {
	entry.logx(ErrorLevel, fmt.Sprintf(format, v...))
}

func (entry *Entry) Errorln(v ...interface{}) {
	entry.logx(ErrorLevel, v...)
}

func (entry *Entry) Warnf(format string, v ...interface{}) {
	entry.logx(WarnLevel, fmt.Sprintf(format, v...))
}

func (entry *Entry) Warnln(v ...interface{}) {
	entry.logx(WarnLevel, v...)
}

func (entry *Entry) Infof(format string, v ...interface{}) {
	entry.logx(InfoLevl, fmt.Sprintf(format, v...))
}

func (entry *Entry) Infoln(v ...interface{}) {
	entry.logx(InfoLevl, v...)
}

func (entry *Entry) Debugf(format string, v ...interface{}) {
	entry.logx(DebugLevl, fmt.Sprintf(format, v...))
}

func (entry *Entry) Debugln(v ...interface{}) {
	entry.logx(DebugLevl, v...)
}

func (entry *Entry) Printf(format string, v ...interface{}) {
	entry.logx(DisableLevel, fmt.Sprintf(format, v...))
}

func (entry *Entry) Println(v ...interface{}) {
	entry.logx(DisableLevel, v...)
}

func (entry *Entry) logx(level Level, v ...interface{}) {
	if entry.logger.Level < level || len(v) == 0 {
		return
	}

	ctx := newCtx(entry.logger, v)
	if entry.logger.Before != nil {
		entry.logger.Before(ctx)
	}
	if entry.logger.After != nil {
		defer entry.logger.After(ctx)
	}

	buf := bytes.Buffer{}
	if entry.logger.Color.ColorMode != Normal {
		if isTerminal(entry.logger.Out) {
			meta := entry.logger.Color.parseMeta(level)
			buf.WriteString(meta.cl(meta.LvName))
		}
	}
	if entry.logger.TimeFormat != "" {
		buf.WriteString(time.Now().Format(entry.logger.TimeFormat))
	}
	var (
		byts []byte
		err  error
	)
	if len(v) > 1 {
		byts, err = entry.logger.Formatter.Format(v)
	} else {
		byts, err = entry.logger.Formatter.Format(v[0])
	}
	if err != nil {
		ctx.Err = err
		return
	}
	if entry.logger.Formatter.NewLine() {
		buf.WriteString("\n")
	}
	buf.Write(byts)
	entry.lg.Output(4, buf.String())
}
