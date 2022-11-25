package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个不转义相应内容的safe模板函数如下：
func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.tmpl")
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})

	router.Run()
}
