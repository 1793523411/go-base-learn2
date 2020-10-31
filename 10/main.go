package main

import (
	"log"
	"os"
)

// 无论是软件开发的调试阶段还是软件上线之后的运行阶段，日志一直都是非常重要的一个环节，我们也应该养成在程序中记录日志的好习惯

// log包定义了Logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用

// func init() {
// 	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("open log file failed, err:", err)
// 		return
// 	}
// 	log.SetOutput(logFile)
// 	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
// }

func main() {
	// log.Println("这是一条很普通的日志。")
	// v := "很普通的"
	// log.Printf("这是一条%s日志。\n", v)
	// log.Fatalln("这是一条会触发fatal的日志。")
	// log.Panicln("这是一条会触发panic的日志。")
	// logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic

	//* 默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法,log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置

	// log标准库提供了如下的flag选项，它们是一系列定义好的常量。
	// const (
	// 	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	// 	Ldate         = 1 << iota     // 日期：2009/01/23
	// 	Ltime                         // 时间：01:23:23
	// 	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	// 	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	// 	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	// 	LUTC                          // 使用UTC时间
	// 	LstdFlags     = Ldate | Ltime // 标准logger的初始值
	// )

	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("这是一条很普通的日志。")

	// log标准库中还提供了关于日志信息前缀的两个方法,其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("这是一条很普通的日志。")
	// log.SetPrefix("[小王子]")
	// log.Println("这是一条很普通的日志。")

	// SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出

	// 如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中
	// log.Println("这是一条很普通的日志。")
	// log.SetPrefix("[小王子]")
	// log.Println("这是一条很普通的日志。")

	// log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下
	// func New(out io.Writer, prefix string, flag int) *Logger

	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.LUTC)
	logger.Println("这是自定义的logger记录的日志。")

	// Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等
}
