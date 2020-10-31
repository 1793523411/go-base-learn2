package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie") // 获取Cookie
		if err != nil {
			cookie = "NotSet"
			// 设置Cookie
			// func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, sameSite http.SameSite, secure, httpOnly bool) {
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", 0, true, false)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
