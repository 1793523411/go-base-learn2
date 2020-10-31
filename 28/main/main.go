package main

import (
	log "github.com/sirupsen/logrus"
)

// 创建一个新的logger实例。可以创建任意多个。
// var log = logrus.New()

func main() {

	// 设置日志输出为os.Stdout
	// log.Out = os.Stdout

	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	// log.WithFields(logrus.Fields{
	// 	"animal": "dog",
	// 	"size":   10,
	// }).Info("一群舔狗出现了。")

	// // 会记录info及以上级别 (warn, error, fatal, panic)
	log.SetLevel(log.WarnLevel)

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	log.Fatal("Bye.")
	// 记完日志后会调用 panic()
	log.Panic("I'm bailing.")

	// // Logrus鼓励通过日志字段进行谨慎的结构化日志记录，而不是冗长的、不可解析的错误消息。
	// log.WithFields(log.Fields{
	// 	"event": event,
	// 	"topic": topic,
	// 	"key":   key,
	// }).Fatal("Failed to send event")

	//!默认字段
	// 通常，将一些字段始终附加到应用程序的全部或部分的日志语句中会很有帮助。例如，你可能希望始终在请求的上下文中记录request_id和user_ip
	// requestLogger := log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})
	// requestLogger.Info("something happened on that request")
	// requestLogger.Warn("something not great happened")

	//!日志条目
	// 除了使用WithField或WithFields添加的字段外，一些字段会自动添加到所有日志记录事中:

	// time：记录日志时的时间戳
	// msg：记录的日志信息
	// level：记录的日志级别

	// 	logrus内置以下两种日志格式化程序：
	// logrus.TextFormatter logrus.JSONFormatter
	// 还支持一些第三方的格式化程序，详见项目首页

	//!记录函数名
	log.SetReportCaller(true)

	//!线程安全
	// 默认的logger在并发写的时候是被mutex保护的，比如当同时调用hook和写log时mutex就会被请求，有另外一种情况，文件是以appending mode打开的， 此时的并发操作就是安全的，可以用logger.SetNoLock()来关闭它

}
