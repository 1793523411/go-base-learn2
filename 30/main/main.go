package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//!Gin中间件
// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	//!为全局路由注册
	r.Use(StatCost())
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	//*定义一个不转义相应内容的safe模板函数如下
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})

	r.Static("/static", "./static")

	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreXML", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})

	// r.GET("/someProtoBuf", func(c *gin.Context) {
	// 	reps := []int64{int64(1), int64(2)}
	// 	label := "test"
	// 	// protobuf 的具体定义写在 testdata/protoexample 文件中。
	// 	data := &protoexample.Test{
	// 		Label: &label,
	// 		Reps:  reps,
	// 	}
	// 	// 请注意，数据在响应中变为二进制数据
	// 	// 将输出被 protoexample.Test protobuf 序列化了的数据
	// 	c.ProtoBuf(http.StatusOK, data)
	// })

	//!获取querystring参数
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//!获取form参数
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//!获取path参数
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("/home/ygjzs/goproject/src/base/30/main/static/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	//!fail
	r.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/home/ygjzs/goproject/src/base/30/main/static/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.ygjie.icu/")
	})
	r.GET("/test0", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	// 	一个可以匹配所有请求方法的Any方法如下：
	// r.Any("/test", func(c *gin.Context) {...})

	r.LoadHTMLFiles("./views/404.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	// //!路由组,一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别
	// userGroup := r.Group("/user")
	// {
	// 	userGroup.GET("/index", func(c *gin.Context) {...})
	// 	userGroup.GET("/login", func(c *gin.Context) {...})
	// 	userGroup.POST("/login", func(c *gin.Context) {...})

	// }
	// shopGroup := r.Group("/shop")
	// {
	// 	shopGroup.GET("/index", func(c *gin.Context) {...})
	// 	shopGroup.GET("/cart", func(c *gin.Context) {...})
	// 	shopGroup.POST("/checkout", func(c *gin.Context) {...})
	// }

	// shopGroup := r.Group("/shop")
	// {
	// 	shopGroup.GET("/index", func(c *gin.Context) {...})
	// 	shopGroup.GET("/cart", func(c *gin.Context) {...})
	// 	shopGroup.POST("/checkout", func(c *gin.Context) {...})
	// 	// 嵌套路由组
	// 	xx := shopGroup.Group("xx")
	// 	xx.GET("/oo", func(c *gin.Context) {...})
	// }
	//* 通常我们将路由分组用在划分业务逻辑或划分API版本时

	//! Gin框架中的路由使用的是httprouter这个库。其基本原理就是构造一个路由地址的前缀树。

	r.GET("/midware", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//!为路由组注册中间件有以下两种写法。
	// 	shopGroup := r.Group("/shop", StatCost())
	// {
	//     shopGroup.GET("/index", func(c *gin.Context) {...})
	//     ...
	// }

	// shopGroup := r.Group("/shop")
	// shopGroup.Use(StatCost())
	// {
	//     shopGroup.GET("/index", func(c *gin.Context) {...})
	//     ...
	// }
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
