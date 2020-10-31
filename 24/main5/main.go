package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
// sync.Once只有一个Do方法，其签名如下：
// func (o *Once) Do(f func()) {}
// 备注：如果要执行的函数f需要传递参数就需要搭配闭包来使用。

//!加载配置文件示例
// 延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践

// var icons map[string]image.Image

// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }

// // Icon 被多个goroutine调用时不是并发安全的
// func Icon(name string) image.Image {
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	return icons[name]
// }

//?使用sync.Once改造的示例代码如下

// var icons map[string]image.Image

// var loadIconsOnce sync.Once

// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }

// // Icon 是并发安全的
// func Icon(name string) image.Image {
// 	loadIconsOnce.Do(loadIcons)
// 	return icons[name]
// }

//!Go语言中内置的map不是并发安全的
// var m = make(map[string]int)

var m = sync.Map{}

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

func main() {
	// wg := sync.WaitGroup{}
	// ?开启少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误,像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法
	// for i := 0; i < 20; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		key := strconv.Itoa(n)
	// 		set(key, n)
	// 		fmt.Printf("k=:%v,v:=%v\n", key, get(key))
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
