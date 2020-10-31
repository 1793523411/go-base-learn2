package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// !Go语言标准库之http/template,html/template包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用html/template这个包

// Go语言内置了文本模板引擎text/template和用于HTML文档的html/template。它们的作用机制可以简单归纳如下：

//     模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
//     模板文件中使用{{和}}包裹和标识需要传入的数据。
//     传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
//     除{{和}}包裹的内容外，其他内容均不做修改原样输出。

//!模板引擎的使用
//* Go语言模板引擎的使用可以分为三部分：定义模板文件、解析模板文件和模板渲染.

// 使用下面的常用方法去解析模板文件，得到模板对象：

// func (t *Template) Parse(src string) (*Template, error)
// func ParseFiles(filenames ...string) (*Template, error)
// func ParseGlob(pattern string) (*Template, error)

// 当然，你也可以使用func New(name string) *Template函数创建一个名为name的模板，然后对其调用上面的方法去解析模板字符串或模板文件。

// 渲染模板简单来说就是使用数据去填充模板，当然实际上可能会复杂很多。

// func (t *Template) Execute(wr io.Writer, data interface{}) error
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error

//!example
// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	// 解析指定文件生成模板对象
// 	tmpl, err := template.ParseFiles("./hello.tmpl")
// 	if err != nil {
// 		fmt.Println("create template failed, err:", err)
// 		return
// 	}
// 	// 利用给定数据渲染模板，并将结果写入w
// 	tmpl.Execute(w, "沙河小王子")
// }

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	// 解析指定文件生成模板对象
// 	tmpl, err := template.ParseFiles("./hello.tmpl")
// 	if err != nil {
// 		fmt.Println("create template failed, err:", err)
// 		return
// 	}
// 	// 利用给定数据渲染模板，并将结果写入w
// 	user := UserInfo{
// 		Name:   "      小王子  ",
// 		Gender: "男",
// 		Age:    18,
// 	}
// 	tmpl.Execute(w, user)
// }

// pipeline是指产生数据的操作。比如{{.}}、{{.Name}}等。Go的模板语法中支持使用管道符号|链接多个命令，用法和unix下的管道类似：|前面的命令会将运算结果(或返回值)传递给后一个命令的最后一个位置。

// 注意：并不是只有使用了|才是pipeline。Go的模板语法中，pipeline的概念是传递数据，只要能产生数据的，都是pipeline

// 有时候我们在使用模板语法的时候会不可避免的引入一下空格或者换行符，这样模板最终渲染出来的内容可能就和我们想的不一样，这个时候可以使用{{-语法去除模板内容左侧的所有空白符号， 使用-}}去除模板内容右侧的所有空白符号
// 注意：-要紧挨{{和}}，同时与模板值之间需要使用空格分隔

//!条件判断

// {{if pipeline}} T1 {{end}}

// {{if pipeline}} T1 {{else}} T0 {{end}}

// {{if pipeline}} T1 {{else if pipeline}} T0 {{end}}

//!range
// {{range pipeline}} T1 {{end}}
// 如果pipeline的值其长度为0，不会有任何输出

// {{range pipeline}} T1 {{else}} T0 {{end}}
// 如果pipeline的值其长度为0，则会执行T0。

//!with

// {{with pipeline}} T1 {{end}}
// 如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。

// {{with pipeline}} T1 {{else}} T0 {{end}}
// 如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。

//!预定义函数
// 执行模板时，函数从两个函数字典中查找：首先是模板函数字典，然后是全局函数字典。一般不在模板内定义函数，而是使用Funcs方法添加函数到模板里
// 预定义的全局函数如下：
// and
//     函数返回它的第一个empty参数或者最后一个参数；
//     就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
// or
//     返回第一个非empty参数或者最后一个参数；
//     亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
// not
//     返回它的单个参数的布尔值的否定
// len
//     返回它的参数的整数类型长度
// index
//     执行结果为第一个参数以剩下的参数为索引/键指向的值；
//     如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
// print
//     即fmt.Sprint
// printf
//     即fmt.Sprintf
// println
//     即fmt.Sprintln
// html
//     返回与其参数的文本表示形式等效的转义HTML。
//     这个函数在html/template中不可用。
// urlquery
//     以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
//     这个函数在html/template中不可用。
// js
//     返回与其参数的文本表示形式等效的转义JavaScript。
// call
//     执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
//     如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
//     其中Y是函数类型的字段或者字典的值，或者其他类似情况；
//     call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
//     该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
//     如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；

//!比较函数
// eq      如果arg1 == arg2则返回真
// ne      如果arg1 != arg2则返回真
// lt      如果arg1 < arg2则返回真
// le      如果arg1 <= arg2则返回真
// gt      如果arg1 > arg2则返回真
// ge      如果arg1 >= arg2则返回真
// 为了简化多参数相等检测，eq（只有eq）可以接受2个或更多个参数，它会将第一个参数和其余参数依次比较，返回下式的结果：
// {{eq arg1 arg2 arg3}}
// 比较函数只适用于基本类型（或重定义的基本类型，如”type Celsius float32”）。但是，整数和浮点数不能互相比较。

//!自定义函数

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./hello.tmpl")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	// 使用user渲染模板，并将结果写入w
	tmpl.Execute(w, user)
}

//!嵌套template
func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

// 在解析模板时，被嵌套的模板一定要在后面解析，例如上面的示例中t.tmpl模板中嵌套了ul.tmpl，所以ul.tmpl要在t.tmpl后进行解析

//!block

// {{block "name" pipeline}} T1 {{end}}

// block是定义模板{{define "name"}} T1 {{end}}和执行{{template "name" pipeline}}缩写，典型的用法是定义一组根模板，然后通过在其中重新定义块模板进行自定义。

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

// 如果我们的模板名称冲突了，例如不同业务线下都定义了一个index.tmpl模板，我们可以通过下面两种方法来解决。

//     在模板文件开头使用{{define 模板名}}语句显式的为模板命名。
//     可以把模板文件存放在templates文件夹下面的不同目录中，然后使用template.ParseGlob("templates/**/*.tmpl")解析模板

//!修改默认的标识符
// Go标准库的模板引擎使用的花括号{{和}}作为标识，而许多前端框架（如Vue和 AngularJS）也使用{{和}}作为标识符，所以当我们同时使用Go语言模板引擎和以上前端框架时就会出现冲突，这个时候我们需要修改标识符，修改前端的或者修改Go语言的。这里演示如何修改Go语言模板引擎默认的标识符：
// template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl")

//! text/template与html/tempalte的区别
// html/template针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击
func xss(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	jsStr := `<script>alert('嘿嘿嘿')</script>`
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println(err)
	}
}

// 这样我们只需要在模板文件不需要转义的内容后面使用我们定义好的safe函数就可以了

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/tmplDemo", tmplDemo)
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
