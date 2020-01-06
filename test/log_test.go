package test

import (
	"fmt"
	"github.com/lhlyu/logger"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logger.Infoln("春又来 人已去")
	logger.Infof("%s %s", "风烟残", "夕阳晚")
	// 设置日志等级: debug info warn error fatal
	logger.SetLevel("debug")
	logger.Debugln("樱花开 顷刻散 零乱")
	// 设置颜色高亮
	logger.SetColorMode(2)
	logger.Warnln("年光逝 韶华落")
	// 设置日期格式
	logger.SetTimeFormat("2006/01/02 ")
	logger.Errorln("飞絮转 不堪看")
	// 设置内容格式化: TEXT JSON JSON_INDENT XML  XML_INDENT YAML,可以自定义
	logger.SetFormatter(logger.JSON)
	a := struct {
		Name string
		Age  int
	}{
		Name: "路漫漫",
		Age:  3,
	}
	logger.Infoln(a)
	logger.SetFormatter(logger.XML)
	// 添加前置处理器,日志打印前处理
	logger.AddBefore(func(ctx *logger.Ctx) {
		fmt.Println("this is first before handler")
	})
	logger.AddBefore(func(ctx *logger.Ctx) {
		fmt.Println("this is second before handler")
		ctx.Stop() // 后面的前置处理器将不会执行
	})
	// 这个不会打印，因为在前一个处理器时已经stop了
	logger.AddBefore(func(ctx *logger.Ctx) {
		fmt.Println("this is third before handler")
	})
	// 添加后置处理器,日志打印处理
	logger.AddAfter(func(ctx *logger.Ctx) {
		if ctx.Err != nil {
			fmt.Println(ctx.Err.Error())
		}
	})
	// 这个不会打印，因为xml不支持这个类型,错误会被上面的后置处理器打印出来
	logger.Infoln(&a)

	// 新建一个日志器
	lg := logger.New()
	// 设置输出流
	lg.SetOutput(os.Stderr)
	entry := logger.NewEntry(lg)
	entry.Fatalln("渡忘川")
}
