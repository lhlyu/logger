package main

import (
	"github.com/lhlyu/logger"
	"fmt"
)

type A struct {
	B int
	C D
}

type D struct {
	E string
	F int
}

func main() {
	test2()
}

func test1(){
	logger.Reset()
	logger.SetTimeFormat("2006-01-02 15.04.05")
	logger.SetLevel(logger.LevelDebug)
	logger.PNone("--> logger.PNone")
	logger.PWarn("--> logger.PWarn")
	logger.SetColorMod(false) // 关闭颜色打印
	logger.PDebug("--> logger.PDebug")
	logger.SetColorMod(true) // 开启颜色打印
	logger.SetDelims("<<<<", ">>>>")
	logger.PDebugf("--> %s", "logger.PDebugf")
	logger.SetLocation(1) // 打印调用方法的位置
	logger.PInfo("--> logger.PInfo")
	logger.SetLocation(2) // 打印调用文件的位置
	logger.PErrorf("--> %s", "logger.PErrorf")
	logger.Reset() // 重置
	logger.PDebug("--> 这个不会输出，因为重置后日志等级默认是1,调试信息的等级是0")
	logger.PWarnf("--> %s", "logger.PWarnf")
	logger.PJson(A{1, D{"logger.PJson", 2}}) // 格式化打印对象
	logger.PJsonL(logger.LevelError, A{1, D{"logger.PJsonL", 2}})
	logger.PNormal("--> logger.PNormal")
	logger.PFatal("--> logger.PFatal") // 打印致命信息会导致程序退出，后面的程序不会再执行
	logger.PInfo("--> 这个不会输出")
}

func test2(){
	fmt.Println(logger.Red(`logger.Red("普通红色字体")`))
	fmt.Println(logger.Green(`logger.Green("普通绿色字体")`))
	fmt.Println(logger.Yellow(`logger.Yellow("普通黄色字体")`))
	fmt.Println(logger.Blue(`logger.Blue("普通蓝色字体")`))
	fmt.Println(logger.Magenta(`logger.Magenta("普通品红字体")`))
	fmt.Println(logger.Cyan(`logger.Cyan("普通青色字体")`))
	fmt.Println(logger.RedLine(`logger.RedLine("红色下划线字体")`))
	fmt.Println(logger.GreenLine(`logger.GreenLine("绿色下划线字体")`))
	fmt.Println(logger.YellowLine(`logger.YellowLine("黄色下划线字体")`))
	fmt.Println(logger.BlueLine(`logger.BlueLine("蓝色下划线字体")`))
	fmt.Println(logger.MagentaLine(`logger.MagentaLine("品红下划线字体")`))
	fmt.Println(logger.CyanLine(`logger.CyanLine("青色下划线字体")`))
	fmt.Println(logger.RedBlod(`logger.RedBlod("粗体红色字体")`))
	fmt.Println(logger.GreenBlod(`logger.GreenBlod("粗体绿色字体")`))
	fmt.Println(logger.YellowBlod(`logger.YellowBlod("粗体黄色字体")`))
	fmt.Println(logger.BlueBlod(`logger.BlueBlod("粗体蓝色字体")`))
	fmt.Println(logger.MagentaBlod(`logger.MagentaBlod("粗体品红字体")`))
	fmt.Println(logger.CyanBlod(`logger.CyanBlod("粗体青色字体")`))
	fmt.Println(logger.RedBg(`logger.RedBg("普通红色背景字体")`))
	fmt.Println(logger.GreenBg(`logger.GreenBg("普通绿色背景字体")`))
	fmt.Println(logger.YellowBg(`logger.YellowBg("普通黄色背景字体")`))
	fmt.Println(logger.BlueBg(`logger.BlueBg("普通蓝色背景字体")`))
	fmt.Println(logger.MagentaBg(`logger.MagentaBg("普通品红背景字体")`))
	fmt.Println(logger.CyanBg(`logger.CyanBg("普通青色背景字体")`))
	fmt.Println(logger.RedLineBg(`logger.RedLineBg("红色下划线背景字体")`))
	fmt.Println(logger.GreenLineBg(`logger.GreenLineBg("绿色下划线背景字体")`))
	fmt.Println(logger.YellowLineBg(`logger.YellowLineBg("黄色下划线背景字体")`))
	fmt.Println(logger.BlueLineBg(`logger.BlueLineBg("蓝色下划线背景字体")`))
	fmt.Println(logger.MagentaLineBg(`logger.MagentaLineBg("品红下划线背景字体")`))
	fmt.Println(logger.CyanLineBg(`logger.CyanLineBg("青色下划线背景字体")`))
	fmt.Println(logger.RedBlodBg(`logger.RedBlodBg("粗体红色背景字体")`))
	fmt.Println(logger.GreenBlodBg(`logger.GreenBlodBg("粗体绿色背景字体")`))
	fmt.Println(logger.YellowBlodBg(`logger.YellowBlodBg("粗体黄色背景字体")`))
	fmt.Println(logger.BlueBlodBg(`logger.BlueBlodBg("粗体蓝色背景字体")`))
	fmt.Println(logger.MagentaBlodBg(`logger.MagentaBlodBg("粗体品红背景字体")`))
	fmt.Println(logger.CyanBlodBg(`logger.CyanBlodBg("粗体青色背景字体")`))
}