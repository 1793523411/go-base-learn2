package main

import (
	"fmt"
	"unsafe"
)

type NewInt int

//类型别名
// type MyInt = int

type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//Person 结构体
type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

func main() {
	// type byte = uint8
	// type rune = int32
	//类型定义

	var a NewInt
	var b MyInt

	// 结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	// 	只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

	// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型

	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	//!匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//!创建指针类型结构体
	// 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	// Go语言中支持对结构体指针直接使用.来访问结构体的成员
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}

	//!取结构体的地址实例化
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
	// p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

	//*没有初始化的结构体，其成员变量都是对应其类型的零值

	//*使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

	//*初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值

	// 必须初始化结构体的所有字段。
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 该方式不能和键值初始化方式混用。

	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}

	//!结构体内存布局
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	//* 空结构体是不占用空间的。
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0

	//!mian

	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	//!构造函数
	p9 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}

	//!方法和接收者
	// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	// 	函数体
	// }

	// 当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身

	p10 := NewPerson("小王子", 25)
	p10.Dream()
	fmt.Println(p10.age)
	p10.SetAge2(99)
	fmt.Println(p10.age)

	// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

	// 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self

	p11 := NewPerson("小王子", 25)
	fmt.Println(p11.age) // 25
	p11.SetAge(30)
	fmt.Println(p11.age) // 30

	//?什么时候应该使用指针类型接收者
	// 需要修改接收者中的值
	// 接收者是拷贝代价比较大的大对象
	// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

	// 在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法

	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt

	////非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

}
