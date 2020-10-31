package main

import "fmt"

var m = 100

func foo() (int, string) {
	return 10, "Q1mi"
}

func main() {
	// var name string
	// var age int
	// var isOk bool
	// var (
	// 	a string
	// 	b int
	// 	c bool
	// 	d float32
	// )
	// var name2 string = "Q1mi"
	// var age2 int = 18
	// var name3, age3 = "Q1mi", 20

	const pi2 = 3.1415
	const e2 = 2.7182
	const (
		pi = 3.1415
		e  = 2.7182
	)

	// const (
	// 	n1 = 100
	// 	n2
	// 	n3
	// )

	// iota是go语言的常量计数器，只能在常量的表达式中使用。
	// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
	// const (
	// 	n1 = iota //0
	// 	n2        //1
	// 	n3        //2
	// 	n4        //3
	// )

	// const (
	// 	n1 = iota //0
	// 	n2        //1
	// 	_
	// 	n4 //3
	// )

	const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	//!zhe li xu yao zhu yi
	const (
		a, b  = iota + 1, iota + 2 //1,2
		c, d                       //2,3
		e3, f                      //3,4
	)

	fmt.Println(m)
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量
	// 		函数外的每个语句都必须以关键字开始（var、const、func等）
	// :=不能使用在函数外。
	// _多用于占位，表示忽略值
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
	fmt.Println(pi, e)
	fmt.Println(n1, n2, n3, n4)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e3, f)
}
