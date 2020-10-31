package main

import (
	"flag"
	"fmt"
	"time"
)

// Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单

func main() {
	// // flag.Type(flag名, 默认值, 帮助信息)*Type 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义
	// name := flag.String("name", "张三", "姓名")
	// age := flag.Int("age", 18, "年龄")
	// married := flag.Bool("married", false, "婚否")
	// delay := flag.Duration("d", 0, "时间间隔")
	// fmt.Println(*name, *age, *married, delay)

	// if len(os.Args) > 0 {
	// 	for index, arg := range os.Args {
	// 		fmt.Printf("args[%d]=%v\n", index, arg)
	// 	}
	// }

	// flag.TypeVar(Type指针, flag名, 默认值, 帮助信息) 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义
	// var name string
	// var age int
	// var married bool
	// var delay time.Duration
	// flag.StringVar(&name, "name", "张三", "姓名")
	// flag.IntVar(&age, "age", 18, "年龄")
	// flag.BoolVar(&married, "married", false, "婚否")
	// flag.DurationVar(&delay, "d", 0, "时间间隔")

	//通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析

	// -flag xxx （使用空格，一个-符号）
	// --flag xxx （使用空格，两个-符号）
	// -flag=xxx （使用等号，一个-符号）
	// --flag=xxx （使用等号，两个-符号）
	// 布尔类型的参数必须使用等号的方式指定
	// Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。

	// flag.Args()  //返回命令行参数后的其他参数，以[]string类型
	// flag.NArg()  //返回命令行参数后的其他参数个数
	// flag.NFlag() //返回使用的命令行参数个数

	// !完整示例
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

	//*$ ./main 沙河娜扎 --age 28 -married=false -d=1h30m
	//*$ ./main a b c

}

// flag包支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration
