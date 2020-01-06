package logger

import (
	"context"
	"fmt"
	"io"
)

var _entry = New().newEntry()

func SetOutput(output io.Writer) *Entry {
	_entry.logger.SetOutput(output)
	return _entry
}

func SetLevel(lv string) *Entry {
	_entry.logger.SetLevel(lv)
	return _entry
}

func SetFormatter(formatter Formatter) *Entry {
	_entry.logger.SetFormatter(formatter)
	return _entry
}

func SetTimeFormat(timeFormat string) *Entry {
	_entry.logger.SetTimeFormat(timeFormat)
	return _entry
}

func WithContext(context context.Context) *Entry {
	_entry.logger.WithContext(context)
	return _entry
}

func SetColorMode(colorMode ColorMode) *Entry {
	_entry.logger.SetColorMode(colorMode)
	return _entry
}

func AddBefore(befor Before) *Entry {
	_entry.logger.AddBefore(befor)
	return _entry
}

func AddAfter(after After) *Entry {
	_entry.logger.AddAfter(after)
	return _entry
}

func Fatalf(format string, v ...interface{}) {
	_entry.logx(FatalLevel, fmt.Sprintf(format, v...))
}

func Fatalln(v ...interface{}) {
	_entry.logx(FatalLevel, v...)
}

func Errorf(format string, v ...interface{}) {
	_entry.logx(ErrorLevel, fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	_entry.logx(ErrorLevel, v...)
}

func Warnf(format string, v ...interface{}) {
	_entry.logx(WarnLevel, fmt.Sprintf(format, v...))
}

func Warnln(v ...interface{}) {
	_entry.logx(WarnLevel, v...)
}

func Infof(format string, v ...interface{}) {
	_entry.logx(InfoLevl, fmt.Sprintf(format, v...))
}

func Infoln(v ...interface{}) {
	_entry.logx(InfoLevl, v...)
}

func Debugf(format string, v ...interface{}) {
	_entry.logx(DebugLevl, fmt.Sprintf(format, v...))
}

func Debugln(v ...interface{}) {
	_entry.logx(DebugLevl, v...)
}

func Printf(format string, v ...interface{}) {
	_entry.logx(DisableLevel, fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	_entry.logx(DisableLevel, v...)
}
