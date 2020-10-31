package main

import "fmt"

//! 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节,在Go语言中接口（interface）是一种类型，一种抽象的类型

// type Cat struct{}

// func (c Cat) Say() string { return "喵喵喵" }

// type Dog struct{}

// func (d Dog) Say() string { return "汪汪汪" }

// type 接口类型名 interface{
//     方法名1( 参数列表1 ) 返回值列表1
//     方法名2( 参数列表2 ) 返回值列表2
//     …
// }

// 接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
// 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
// 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略

type writer interface {
	Write([]byte) error
}

// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表

type dog struct{}

type cat struct{}

// Sayer 接口
type Sayer interface {
	say()
}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	// //有重复的代码
	// c := Cat{}
	// fmt.Println("猫:", c.Say())
	// d := Dog{}
	// fmt.Println("狗:", d.Say())

	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}
