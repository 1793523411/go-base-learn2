package main

import (
	"fmt"
)

type Person2 struct {
	string
	int
}

// type Address struct {
// 	Province string
// 	City     string
// }

// //User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string
// 	Address //匿名字段
// }

// type Address struct {
// 	Province string
// 	City     string
// }

// //User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string
// 	Address //匿名字段
// }

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// //Student 学生
// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// }

// //Class 班级
// type Class struct {
// 	Title    string
// 	Students []*Student
// }

type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

// func (p *Person) SetDreams(dreams []string) {
// 	p.dreams = dreams
// }

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

// 同样的问题也存在于返回值slice和map的情况

func main() {
	//!结构体的匿名字段
	// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

	// p1 := Person2{
	// 	"小王子",
	// 	18,
	// }
	// fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	// fmt.Println(p1.string, p1.int) //北京 18
	////这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个

	//!嵌套结构体
	// user1 := User{
	// 	Name:   "小王子",
	// 	Gender: "男",
	// 	Address: Address{
	// 		Province: "山东",
	// 		City:     "威海",
	// 	},
	// }
	// fmt.Printf("user1=%#v\n", user1) //user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	// 上面user结构体中嵌套的Address结构体也可以采用匿名字段的方式，
	// var user2 User
	// user2.Name = "小王子"
	// user2.Gender = "男"
	// user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	// user2.City = "威海"                // 匿名字段可以省略
	// fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	// 嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名

	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019"         //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
	fmt.Println(user3)

	//!结构体的“继承”
	// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！

	//!结构体字段的可见性
	// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

	////--------------------------------------------------
	// //!结构体与JSON序列化
	// c := &Class{
	// 	Title:    "101",
	// 	Students: make([]*Student, 0, 200),
	// }
	// for i := 0; i < 10; i++ {
	// 	stu := &Student{
	// 		Name:   fmt.Sprintf("stu%02d", i),
	// 		Gender: "男",
	// 		ID:     i,
	// 	}
	// 	c.Students = append(c.Students, stu)
	// }
	// // fmt.Println(c)
	// //JSON序列化：结构体-->JSON格式的字符串
	// data, err := json.Marshal(c)
	// if err != nil {
	// 	fmt.Println("json marshal failed")
	// 	return
	// }
	// fmt.Printf("json:%s\n", data)
	// //JSON反序列化：JSON格式的字符串-->结构体
	// str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	// c1 := &Class{}
	// err = json.Unmarshal([]byte(str), c1)
	// if err != nil {
	// 	fmt.Println("json unmarshal failed!")
	// 	return
	// }
	// fmt.Printf("%#v\n", c1)

	//!结构体标签（Tag）
	// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。

	// s1 := Student{
	// 	ID:     1,
	// 	Gender: "男",
	// 	name:   "沙河娜扎",
	// }
	// data, err := json.Marshal(s1)
	// if err != nil {
	// 	fmt.Println("json marshal failed!")
	// 	return
	// }
	// fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}

	// 因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意

	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?
	// 正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
}
