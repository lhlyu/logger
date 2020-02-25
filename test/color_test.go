package test

import (
	"fmt"
	"github.com/lhlyu/logger/v3/color"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(color.Blue("春又来 人已去"))
	fmt.Println(color.Red("风烟残 夕阳晚"))
	fmt.Println(color.Yellowf("%s", "樱花开 顷刻散 零乱"))
	// 设置颜色模式:  0 - 没有颜色; 1 - 颜色(默认) ; 2 - 高亮 ; 3 - 反显
	color.SetMode(2)
	fmt.Println(color.Green("年光逝 韶华落"))
	color.SetMode(3)
	fmt.Println(color.White("飞絮转 不堪看"))
	color.SetMode(0)
	fmt.Println(color.Magenta("路漫漫 空梦断 零乱"))
	// 新建一个
	clr := color.NewColor()
	fmt.Println(clr.Cyan("渡忘川 彼岸 忘不掉 人常叹 古井下 月光思念装满 樱花瓣 飞过 风幽怨 水清叹 离伤黯 游丝转 零乱"))
}
