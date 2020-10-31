package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

//!使用模板继承
// Gin框架默认都是使用单模板，如果需要使用block template功能，可以通过"github.com/gin-contrib/multitemplate"库实现，具体示例如下：

// 首先，假设我们项目目录下的templates文件夹下有以下模板文件，其中home.tmpl和index.tmpl继承了base.tmpl

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}
