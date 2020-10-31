package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
	// 数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化
	// var a [3]int
	// var b [4]int
	// // a = b //不可以这样做，因为此时a和b是不同的类型

	// a := [...]int{1: 1, 3: 5}
	// fmt.Println(a)                  // [0 1 0 5]
	// fmt.Printf("type of a:%T\n", a) //type of a:[4]int

	// var a = [...]string{"北京", "上海", "深圳"}
	// // 方法1：for循环遍历
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// // 方法2：for range遍历
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }
	// fmt.Println(a)
	// // fmt.Println(b)

	// a := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	// fmt.Println(a[2][1]) //支持索引取值:重庆

	// a := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// for _, v1 := range a {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s\t", v2)
	// 	}
	// 	fmt.Println()
	// }

	//支持的写法
	// a := [...][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// //不支持多维数组的内层使用...
	// b := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(a)

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
	//!数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的,the Array of diffent size is not the same type
	//![n]*T表示指针数组，*[n]T表示数组指针 。

}
