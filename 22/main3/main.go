package main

import "fmt"

//!多个类型实现同一接口
// Mover 接口
// type Mover interface {
// 	move()
// }

type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

//!一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

//!接口嵌套

// 接口与接口间可以通过嵌套创造出新的接口。
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

//!空接口

// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。

// 使用空接口实现可以接收任意类型的函数参数
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 使用空接口实现可以保存任意值的字典。

//!类型断言
// 空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢
// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值

// 想要判断空接口中的值这个时候就可以使用类型断言,其语法格式：
// x.(T)
// x：表示类型为interface{}的变量
// T：表示断言x可能是的类型。
//*该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

// 因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛
// !关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗

func main() {
	// var x Mover
	// var a = dog{name: "旺财"}
	// var b = car{brand: "保时捷"}
	// x = a
	// x.move()
	// x = b
	// x.move()

	// var x animal
	// x = cat{name: "花花"}
	// x.move()
	// x.say()

	// // 定义一个空接口x
	// var x interface{}
	// s := "Hello 沙河"
	// x = s
	// fmt.Printf("type:%T value:%v\n", x, x)
	// i := 100
	// x = i
	// fmt.Printf("type:%T value:%v\n", x, x)
	// b := true
	// x = b
	// fmt.Printf("type:%T value:%v\n", x, x)

	show(1)
	show("aaa")
	show(true)

	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
