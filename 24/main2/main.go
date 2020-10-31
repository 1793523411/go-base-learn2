package main

import "fmt"

// 有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收,Go语言中提供了单向通道来处理这种情况

// chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
// <-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// // 当通道被关闭时，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值。那如何判断一个通道是否被关闭了呢？
	// //for range从通道循环取值
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// // 开启goroutine将0~100的数发送到ch1中
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		ch1 <- i
	// 	}
	// 	close(ch1)
	// }()
	// // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	// go func() {
	// 	for {
	// 		i, ok := <-ch1 // 通道关闭后再取值ok=false
	// 		if !ok {
	// 			break
	// 		}
	// 		ch2 <- i * i
	// 	}
	// 	close(ch2)
	// }()
	// // 有两种方式在接收值的时候判断该通道是否被关闭，不过我们通常使用的是for range的方式。使用for range遍历通道，当通道被关闭的时候就会退出for range
	// // 在主goroutine中从ch2中接收值打印
	// for i := range ch2 { // 通道关闭后会退出for range循环
	// 	fmt.Println(i)
	// }
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
