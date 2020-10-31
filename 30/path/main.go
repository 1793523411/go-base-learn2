package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 关于模板文件和静态文件的路径，我们需要根据公司/项目的要求进行设置。可以使用下面的函数获取当前执行程序的路径。

func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}

func main() {
	path := getCurrentPath()
	fmt.Println(path)
}
