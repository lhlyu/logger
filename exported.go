package logger

import (
	"context"
	"fmt"
	"io"
)

var _log = New()

func SetOutput(output io.Writer) *Logger {
	_log.SetOutput(output)
	return _log
}

func SetLevel(lv string) *Logger {
	_log.SetLevel(lv)
	return _log
}

func SetFormatter(formatter Formatter) *Logger {
	_log.SetFormatter(formatter)
	return _log
}

func SetTimeFormat(timeFormat string) *Logger {
	_log.SetTimeFormat(timeFormat)
	return _log
}

func WithContext(context context.Context) *Logger {
	_log.WithContext(context)
	return _log
}

func SetColorMode(colorMode ColorMode) *Logger {
	_log.SetColorMode(colorMode)
	return _log
}

func AddBefore(befor Before) *Logger {
	_log.AddBefore(befor)
	return _log
}

func AddAfter(after After) *Logger {
	_log.AddAfter(after)
	return _log
}

func Fatalf(format string, v ...interface{}) {
	_log.logx(FatalLevel, fmt.Sprintf(format, v...))
}

func Fatalln(v ...interface{}) {
	_log.logx(FatalLevel, v...)
}

func Errorf(format string, v ...interface{}) {
	_log.logx(ErrorLevel, fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	_log.logx(ErrorLevel, v...)
}

func Warnf(format string, v ...interface{}) {
	_log.logx(WarnLevel, fmt.Sprintf(format, v...))
}

func Warnln(v ...interface{}) {
	_log.logx(WarnLevel, v...)
}

func Infof(format string, v ...interface{}) {
	_log.logx(InfoLevl, fmt.Sprintf(format, v...))
}

func Infoln(v ...interface{}) {
	_log.logx(InfoLevl, v...)
}

func Debugf(format string, v ...interface{}) {
	_log.logx(DebugLevl, fmt.Sprintf(format, v...))
}

func Debugln(v ...interface{}) {
	_log.logx(DebugLevl, v...)
}

func Printf(format string, v ...interface{}) {
	_log.logx(DisableLevel, fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	_log.logx(DisableLevel, v...)
}
