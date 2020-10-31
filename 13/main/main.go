package main

import (
	"errors"
	"fmt"
	"strings"
)

// Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”
// !调用有返回值的函数时，可以不接收其返回值
//!参数
// 函数的参数中如果相邻变量的类型相同，则可以省略类型,可变参数是指函数的参数数量不固定。
// o语言中的可变参数通过在参数名后加...来标识,可变参数通常要作为函数的最后一个参数,固定参数搭配可变参数使用时，可变参数要放在固定参数的后面,本质上，函数的可变参数是通过切片来实现的。
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//!返回值
// // Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来
// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	sub := x - y
// 	return sum, sub
// }

// // 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
// func calc2(x, y int) (sum, sub int) {
// 	sum = x + y
// 	sub = x - y
// 	return
// }

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return nil
}

//!函数进阶

//!变量作用域
//全局变量,,局部变量

//!函数类型与变量
// 我们可以使用type关键字来定义一个函数类型，具体格式如下

// 上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值
// 简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型

// func add(x, y int) int {
// 	return x + y
// }

func sub(x, y int) int {
	return x - y
}

type calculation func(int, int) int

//!高阶函数
// 高阶函数分为函数作为参数和函数作为返回值两部分

func add(x, y int) int {
	return x + y
}

// func calc(x, y int, op func(int, int) int) int {
// 	return op(x, y)
// }

//!函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//!匿名函数和闭包
// 函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下

// func(参数)(返回值){
//     函数体
// }

func niming() {
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

// 匿名函数多用于实现回调函数和闭包

//!闭包

// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效

//!闭包进阶示例

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// HasSuffix 判断字符串 s 是否以 suffix 结尾：
// strings.HasSuffix(s, suffix string) bool

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}

// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return add, sub
// }

// 闭包其实并不复杂，只要牢记闭包=函数+引用环境

//!defer语句

// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行

// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等

//!defer执行时机

// 在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前

//!defer经典案例

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//!defer面试题

//!defer注册要延迟执行的函数时该函数所有的参数都需要确定其值

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//!内置函数介绍

// close 	主要用来关闭channel
// len 	用来求长度，比如string、array、slice、map、channel
// new 	用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make 	用来分配内存，主要用来分配引用类型，比如chan、map、slice
// append 	用来追加元素到数组、slice中
// panic和recover 	用来做错误处理

//!panic/recover

// Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效

// func funcA() {
// 	fmt.Println("func A")
// }

// func funcB() {
// 	panic("panic in B")
// }

// func funcC() {
// 	fmt.Println("func C")
// }

// 程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。这个时候我们就可以通过recover将程序恢复回来，继续往后执行

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。

func main() {
	// ret1 := intSum2()
	// ret2 := intSum2(10)
	// ret3 := intSum2(10, 20)
	// ret4 := intSum2(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	// var c calculation               // 声明一个calculation类型的变量c
	// c = add                         // 把add赋值给c
	// fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	// fmt.Println(c(1, 2))            // 像调用add一样调用c

	// f := sub                        // 将函数add赋值给变量f1
	// fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	// fmt.Println(f(10, 20))          // 像调用add一样调用f

	// ret2 := calc(10, 20, add)
	// fmt.Println(ret2) //30

	// var f = adder()
	// fmt.Println(f(10)) //10
	// fmt.Println(f(20)) //30
	// fmt.Println(f(30)) //60

	// f1 := adder()
	// fmt.Println(f1(40)) //40
	// fmt.Println(f1(50)) //90

	// var f = adder2(10)
	// fmt.Println(f(10)) //20
	// fmt.Println(f(20)) //40
	// fmt.Println(f(30)) //70

	// f1 := adder2(20)
	// fmt.Println(f1(40)) //60
	// fmt.Println(f1(50)) //110

	// jpgFunc := makeSuffixFunc(".jpg")
	// txtFunc := makeSuffixFunc(".txt")
	// fmt.Println(jpgFunc("test"))     //test.jpg
	// fmt.Println(txtFunc("test"))     //test.txt
	// fmt.Println(txtFunc("test.txt")) //test.txt
	// fmt.Println(txtFunc("test.jpg")) //test.txt

	// f1, f2 := calc(10)
	// fmt.Println(f1(1), f2(2)) //11 9
	// fmt.Println(f1(3), f2(4)) //12 8
	// fmt.Println(f1(5), f2(6)) //13 7

	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")

	// //分析过程：定义x = 5，然后将返回值int赋值为x也就是5，此时defer函数执行，x变为6，但此时int值并不受影响（可以理解为两块内存地址），再执行RET所以返回值为5
	// fmt.Println(f1())	// 5
	// //!分析过程：调用f2函数，将y返回值x 赋值为5，此时y执行defer函数将x 变为6，同时返回值变为6（可以理解为同一块内存地址），再执行RET所以返回值为6
	// fmt.Println(f2())	// 6
	// //分析过程：定义x = 5，然后将返回值y赋值为x也就是5，此时defer函数执行，x变为6，但此时y值并不受影响（可以理解为两块内存地址），再执行RET所以返回值为5
	// fmt.Println(f3())	// 5
	// //!分析过程：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值，因此在执行到defer时，x = 0。执行到return时，将 返回值x赋值为5，此时defer函数中x 变为1，但对返回值x没有影响（可以理解为两块内存地址）， 再执行RET所以返回值为5
	// fmt.Println(f4())	// 5

	// x := 1
	// y := 2
	// defer calc("AA", x, calc("A", x, y)) ////在执行到此处，因为y参数没有确定，因此首先会调用calc函数，运算得出y值。
	// x = 10
	// defer calc("BB", x, calc("B", x, y)) ////在执行到此处，因为y参数没有确定，因此首先会调用calc函数，运算得出y值。
	// y = 20

	// 	A 1 2 3
	// B 10 2 12
	// BB 10 12 22
	// AA 1 3 4

	funcA()
	funcB()
	funcC()
}
