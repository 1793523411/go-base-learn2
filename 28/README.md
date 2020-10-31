**Logrus是Go（golang）的结构化logger，与标准库logger完全API兼容。**

它有以下特点：

+ 完全兼容标准日志库，拥有七种日志级别：Trace, Debug, Info, Warning, Error, Fataland Panic。
+   可扩展的Hook机制，允许使用者通过Hook的方式将日志分发到任意地方，如本地文件系统，logstash，elasticsearch或者mq等，或者通过Hook定义日志内容和格式等
+    可选的日志输出格式，内置了两种日志格式JSONFormater和TextFormatter，还可以自定义日志格式
+   Field机制，通过Filed机制进行结构化的日志记录
+   线程安全

`go get github.com/sirupsen/logrus`
