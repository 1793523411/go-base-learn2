```
类型 	格式 	作用
测试函数 	函数名前缀为Test 	测试程序的一些逻辑行为是否正确
基准函数 	函数名前缀为Benchmark 	测试函数的性能
示例函数 	函数名前缀为Example 	为文档提供示例文档

```

## 测试函数

```bash
go test

go test -v #查看测试函数名称和运行时间

go test -v -run="More" #添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执xing

go test -cover  #查看测试覆盖率

go test -cover -coverprofile=c.out #将覆盖率相关的记录信息输出到一个文件

go tool cover -html=c.out # 使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告
```

## 基准函数

```bash
go test -bench=Split # 基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试

go test -bench=Split -benchmem #添加-benchmem参数，来获得内存分配的统计数据

go test -bench=. #运行基准测试

go test -bench=Fib40 -benchtime=20s #使用-benchtime标志增加最小基准时间，以产生更准确的结果

```

## 示例函数

示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联
```shell
go test -run Example # 函数只要包含了// Output:也是可以通过go test运行的可执行测试
```