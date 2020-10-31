package main

import (
	"fmt"
	"time"
)

//!worker pool（goroutine池）
// 在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

//!select多路复用
// *select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：
// select{
//     case <-ch1:
//         ...
//     case data := <-ch2:
//         ...
//     case ch3<-data:
//         ...
//     default:
//         默认操作
// }

// 使用select语句能提高代码的可读性。

//     可处理一个或多个channel的发送/接收操作。
//     如果多个case同时满足，select会随机选择一个。
//     对于没有case的select{}会一直等待，可用于阻塞main函数。

func main() {
	// jobs := make(chan int, 100)
	// results := make(chan int, 100)
	// // 开启3个goroutine
	// for w := 1; w <= 3; w++ {
	// 	go worker(w, jobs, results)
	// }
	// // 5个任务
	// for j := 1; j <= 5; j++ {
	// 	jobs <- j
	// }
	// close(jobs)
	// // 输出结果
	// for a := 1; a <= 5; a++ {
	// 	fmt.Println(<-results)
	// }

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(i) //1 3 5 ...
			fmt.Println(x) //0 2 4 6...
		case ch <- i:
		}
	}
}
