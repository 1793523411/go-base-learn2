package main

import "fmt"

//!值接收者和指针接收者实现接口的区别

// type Mover interface {
// 	move()
// }

// type dog struct{}

// // func (d dog) move() {
// // 	fmt.Println("狗会动")
// // }

// func (d *dog) move() {
// 	fmt.Println("狗会动")
// }

////--------------------------------------------

//!类型与接口的关系
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	// var x Mover
	// var wangcai = dog{} // 旺财是dog类型
	// x = wangcai         // x可以接收dog类型
	// var fugui = &dog{}  // 富贵是*dog类型
	// x = fugui           // x可以接收*dog类型
	// x.move()
	// // 使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui

	// var x Mover
	// // var wangcai = dog{} // 旺财是dog类型
	// // x = wangcai         // x不可以接收dog类型
	// var fugui = &dog{} // 富贵是*dog类型
	// x = fugui          // x可以接收*dog类型
	// x.move()
	// // !此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值

	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}
