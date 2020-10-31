package main

import (
	"fmt"
	"sync"
)

// Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴

// func hello() {
// 	fmt.Println("Hello Goroutine!")
// }

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

//! goroutine与线程
// OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的goroutine也是可以的。
// GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

// Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

//!channel
// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
// 虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
// Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制

// var ch chan int
// fmt.Println(ch) // <nil>

// 通道是引用类型，通道类型的空值是nil,声明的通道后需要使用make函数初始化之后才能使用
// ch4 := make(chan int)
// ch5 := make(chan bool)
// ch6 := make(chan []int)

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	// go hello() // 启动另外一个goroutine去执行hello函数
	// fmt.Println("main goroutine done!")
	// time.Sleep(time.Second)
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) // 启动一个goroutine就登记+1
	// 	go hello(i)
	// }
	// wg.Wait() // 等待所有登记的goroutine都结束

	// 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务
	// runtime.GOMAXPROCS(2)
	// go a()
	// go b()
	// time.Sleep(time.Second)

	// ch := make(chan int)
	// // 我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送
	// ch <- 10
	// // 代码会阻塞在ch <- 10这一行代码形成死锁
	// ch <- 20
	// x := <-ch
	// fmt.Println(x)
	// // fmt.Println(<-ch)
	// close(ch)
	// // 通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的

	// // 对一个关闭的通道再发送值就会导致panic。
	// // 对一个关闭的通道进行接收会一直获取值直到通道为空。
	// // 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	// // 关闭一个已经关闭的通道会导致panic。

	// ch := make(chan int)
	// go recv(ch) // 启用goroutine从通道接收值
	// ch <- 10
	// fmt.Println("发送成功")
	// // 	无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	// // 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")

}
