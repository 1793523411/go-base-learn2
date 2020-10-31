package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 沙河！")
}

// 要管理服务端的行为，可以创建一个自定义的Server：

// s := &http.Server{
// 	Addr:           ":8080",
// 	Handler:        myHandler,
// 	ReadTimeout:    10 * time.Second,
// 	WriteTimeout:   10 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }
// log.Fatal(s.ListenAndServe())

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
