# logger
日志打印工具

### 安装

> go get github.com/lhlyu/logger

### 日志等级

```
DEBUG   debug信息  0   不会打印文件和行号    默认等级
INFO    普通信息   1   不会打印文件和行号
CONFIG  配置信息   2   不会打印文件和行号
SIGN    标记信息   3   会打印文件和行号
ERROR   错误信息   4   会打印文件和行号  
FATAL   致命信息   5   会打印文件和行号并且退出程序
```

### 使用 

- 直接使用

```go
import "github.com/lhlyu/logger"

func main() {
	logger.Debug("Lv.0")
	logger.Info("Lv.1")
	logger.Config("Lv.2")
	logger.Sign("Lv.3")
	logger.Error("Lv.4")
	logger.Fatal("Lv.5")
}
```
!["效果1"](./img/console1.jpg)

### 其他方法
```
logger.XXXX(v ...interface{})                     // 打印
logger.XXXXf(format string,v ...interface{})      // 格式化打印
logger.NewLogger(lv int, fldir string) *Logger    // New
logger.SetLogger(logger *Logger)                  // 设置自己的日志管理器
logger.SetColor(open int)                         // 设置颜色打印，默认(open=0)是开启
logger.SetLevel(level int)                        // 设置日志等级，默认是 0 
logger.SetAbs(abs int)                            // 设置打印文件位置是否使用绝对路径，默认0，不使用 
```


### 注意

- 有的终端是不支持色彩打印的，所以如果需要设置不打印色彩提示，可以取消色彩
- 文本日志是默认不开启色彩

```go
logger.SetColor(1)  // 关闭色彩
logger.SetColor(0)  // 开启色彩，这是默认设置
```


