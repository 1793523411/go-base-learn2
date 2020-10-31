package main

import (
	"fmt"
	"sort"
)

func main() {
	// 因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性,切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容,切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

	// 声明切片类型
	// var a []string              //声明一个字符串切片
	// var b = []int{}             //声明一个整型切片并初始化
	// var c = []bool{false, true} //声明一个布尔切片并初始化
	// // var d = []bool{false, true} //声明一个布尔切片并初始化
	// fmt.Println(a)        //[]
	// fmt.Println(b)        //[]
	// fmt.Println(c)        //[false true]
	// fmt.Println(a == nil) //true
	// fmt.Println(b == nil) //false
	// fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量
	// 切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。
	// 切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的low和high表示一个索引范围（左包含，又不包含），也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。

	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3] // s := a[low:high]
	// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// a[2:] // 等同于 a[2:len(a)]
	// a[:3] // 等同于 a[0:3]
	// a[:]  // 等同于 a[0:len(a)]

	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3] // s := a[low:high]
	// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	// fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	// a := [5]int{1, 2, 3, 4, 5}
	// t := a[1:3:5]
	// fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	// 我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数

	// a := make([]int, 2, 10)
	// fmt.Println(a)      //[0 0]
	// fmt.Println(len(a)) //2
	// fmt.Println(cap(a)) //10

	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil

	// var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	// s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	// s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil

	// 拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	// s1 := make([]int, 3) //[0 0 0]
	// s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	// s2[0] = 100
	// fmt.Println(s1) //[100 0 0]
	// fmt.Println(s2) //[100 0 0]

	// Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中

	// copy()复制切片
	// a := []int{1, 2, 3, 4, 5}
	// c := make([]int, 5, 5)
	// copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	// fmt.Println(a) //[1 2 3 4 5]
	// fmt.Println(c) //[1 2 3 4 5]
	// c[0] = 1000
	// fmt.Println(a) //[1 2 3 4 5]
	// fmt.Println(c) //[1000 2 3 4 5]

	// s := []int{1, 3, 5}

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }

	// for index, value := range s {
	// 	fmt.Println(index, value)
	// }

	// var s []int
	// s = append(s, 1)       // [1]
	// s = append(s, 2, 3, 4) // [1 2 3 4]
	// s2 := []int{5, 6, 7}
	// s = append(s, s2...) // [1 2 3 4 5 6 7]
	// fmt.Println(s)

	// 通过var声明的零值切片可以在append()函数直接使用，无需初始化。

	// s := []int{} // 没有必要初始化
	// s = append(s, 1, 2, 3)

	// var s = make([]int) // 没有必要初始化
	// s = append(s, 1, 2, 3)

	// var s []int
	// fmt.Println(s == nil, s)
	// s = append(s, 1, 2, 3)
	// fmt.Println(s == nil, s)

	// 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值

	//append()添加元素和切片扩容
	// var numSlice []int
	// for i := 0; i < 10; i++ {
	// 	numSlice = append(numSlice, i)
	// 	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	// }
	// 结果可以看出:append()函数将元素追加到切片的最后并返回该切片。切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍

	// var citySlice []string
	// // 追加一个元素
	// citySlice = append(citySlice, "北京")
	// // 追加多个元素
	// citySlice = append(citySlice, "上海", "广州", "深圳")
	// // 追加切片
	// a := []string{"成都", "重庆"}
	// citySlice = append(citySlice, a...)
	// fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]。

	// 切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样

	// Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	// 从切片中删除元素
	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// // 要删除索引为2的元素
	// a = append(a[:2], a[3:]...)
	// fmt.Println(a) //[30 31 33 34 35 36 37]
	// // 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)

	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)
	// 初始化的时候就有5个空元素，总容量其实没什么关系，因为循环添加了10个元素，原来的容量不足了，会自动扩容。所以一共应该会有15个元素

	var a = [...]int{3, 7, 8, 9, 1}
	// fmt.Println(a[:]) //将数组切片
	//sort包内部实现了内部数据类型的排序
	//升序排列
	sort.Ints(a[:])
	fmt.Println(a)
	//降序排列
	// func Reverse(data Interface) Interface,Reverse包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列。
	// Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)
}
