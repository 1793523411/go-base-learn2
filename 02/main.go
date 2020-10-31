package main

import (
	"fmt"
	"math"
)

// Go 语言的字符有以下两种：

//     uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
//     rune类型，代表一个 UTF-8字符。

// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
// Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	// 因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串
	// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
// 强制类型转换的基本语法T()
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func count(s string) int {
	count := 0
	count2 := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	for _, a := range s {
		fmt.Println("+++", a)
		count2++
	}
	return int((count - count2) / 2)
}

func count2(s string) int {
	temp := []rune(s)
	count := 0
	for _, v := range temp {
		if v > 256 {
			count++
			fmt.Println(string(v))
		}
	}
	return count
}
func main() {
	// 整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
	// 其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型
	// 在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。
	// 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用int来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint

	// Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%d \n", b) // 77
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%xd\n", c) // ff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64

	fmt.Printf("%f\n", math.Pi) //3.141593
	fmt.Printf("%.2f\n", math.Pi)

	// 复数:complex64和complex128
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
	s1 := "hello"
	s2 := "你好"

	fmt.Println(s1, s2)

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	s3 :=
		`第一行
第二行
第三行
		`
	fmt.Println(s3)

	//! Common operations on strings
	// len(str) 	求长度
	// +或fmt.Sprintf 	拼接字符串
	// strings.Split 	分割
	// strings.contains 	判断是否包含
	// strings.HasPrefix,strings.HasSuffix 	前缀/后缀判断
	// strings.Index(),strings.LastIndex() 	子串出现的位置
	// strings.Join(a[]string, sep string) 	join操作

	// 组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来

	// var a2 := '中'
	// var b2 := 'x'

	// fmt.Println(a2,b2)
	traversalString()
	changeString()
	sqrtDemo()
	fmt.Println(count("hello沙河小王子"))
	fmt.Println(count2("hello沙河小王子"))
}
