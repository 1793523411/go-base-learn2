package main

//!Hooks
// 你可以添加日志级别的钩子（Hook）。例如，向异常跟踪服务发送Error、Fatal和Panic、信息到StatsD或同时将日志发送到多个位置，例如syslog。
// Logrus配有内置钩子。在init中添加这些内置钩子或你自定义的钩子

import (
	log "github.com/sirupsen/logrus" // the package is named "airbrake"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
)

func init() {

	// Use the Airbrake hook to report errors that have Error severity or above to
	// an exception tracker. You can create custom hooks, see the Hooks section.
	// ?log.AddHook(airbrake.NewHook(123, "xyz", "production"))

	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
}

func main() {

}
