package main

import "fmt"

// Go语言设计模式之函数式选项模式

type Option struct {
	A string
	B string
	C int
}

func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

type OptionFunc func(*Option)

func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}

var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)

func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}

func main() {
	// res := newOption("a", "b", 3)
	// fmt.Printf("%#v", res)

	x := newOption("nazha", "小王子", 10)
	fmt.Println(x)
	x = newOption2()
	fmt.Println(x)
	x = newOption2(
		WithA("沙河娜扎"),
		WithC(250),
	)
	fmt.Println(x)
}

// 这样一个使用函数式选项设计模式的构造函数就实现了。这样默认值也有了，以后再要为Option添加新的字段也不会影响之前的代码
