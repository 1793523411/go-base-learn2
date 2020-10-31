package main

import (
	"fmt"
	"reflect"
)

//! 结构体反射
// 与结构体相关的方法
// 任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息

// reflect.Type中与获取结构体成员相关的的方法如下表所示。
// 方法 	说明
// *Field(i int) StructField 	根据索引，返回索引对应的结构体字段的信息。
// *NumField() int 	返回结构体成员字段数量。
// *FieldByName(name string) (StructField, bool) 	根据给定字符串返回字符串对应的结构体字段的信息。
// *FieldByIndex(index []int) StructField 	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
// *FieldByNameFunc(match func(string) bool) (StructField,bool) 	根据传入的匹配函数匹配需要的字段。
// *NumMethod() int 	返回该类型的方法集中方法的数目
// *Method(int) Method 	返回该类型方法集中的第i个方法
// *MethodByName(string)(Method, bool) 	根据方法名返回该类型方法集中的方法

// type StructField struct {
//     // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
//     // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
//     Name    string
//     PkgPath string
//     Type      Type      // 字段的类型
//     Tag       StructTag // 字段的标签
//     Offset    uintptr   // 字段在结构体中的字节偏移量
//     Index     []int     // 用于Type.FieldByIndex时的索引切片
//     Anonymous bool      // 是否匿名字段
// }

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

// 反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。

//     基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
//     大量使用反射的代码通常难以理解。
//     反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	// t := reflect.TypeOf(stu1)
	// fmt.Println(t.Name(), t.Kind()) // student struct
	// // 通过for循环遍历结构体的所有字段信息
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	// }

	// // 通过字段名获取指定结构体字段信息
	// if scoreField, ok := t.FieldByName("Score"); ok {
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	// }

	printMethod(stu1)
}
