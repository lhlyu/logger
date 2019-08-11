package logger

import "os"

// 这样命名有利于 提示
const (
	LevelNone  = 5 // 无
	LevelFatal = 4 // 致命信息 红色
	LevelError = 3 // 错误信息 品红色
	LevelWarn  = 2 // 警告信息 黄色
	LevelInfo  = 1 // 普通信息 蓝色
	LevelDebug = 0 // 调试信息 绿色
)

type levelMetadata struct {
	Level int
	Txt   string
	ColorHanler func(s string) string
	NextHandler func()
}

var exit = func() {
	os.Exit(0)
}

var levelMap = map[int]levelMetadata{
	LevelNone:  levelMetadata{5,"",nil,nil},
	LevelFatal: levelMetadata{4,"FATAL",RedBlod,exit},
	LevelError: levelMetadata{3,"ERROR",MagentaBlod,nil},
	LevelWarn:  levelMetadata{2,"WARN ",YellowBlod,nil},
	LevelInfo:  levelMetadata{1,"INFO ",BlueBlod,nil},
	LevelDebug: levelMetadata{0,"DEBUG",GreenBlod,nil},
}

func getLevelMetadata(level int)levelMetadata{
	if v,ok := levelMap[level];ok{
		return v
	}
	return levelMetadata{5,"",nil,nil}
}