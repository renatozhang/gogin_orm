package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//静态文件：
//html页面上用到的样式文件.css js文件图片，
func main() {
	router := gin.Default()

	router.Static("/static/", "./static")

	// gin框架中给模板添加自定义函数
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// router.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl"")
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.tmpl",
			gin.H{
				"title": "posts/index.tmpl",
			})
	})
	router.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.tmpl",
			gin.H{
				"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
			})
	})

	// 从网上下载的模板
	router.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.Run()
}
