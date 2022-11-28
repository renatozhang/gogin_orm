package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取form表单提交的数据
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	// login post
	r.POST("/login", func(ctx *gin.Context) {
		//获取form表单提交的数据
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password") //渠道就返回值，娶不到返回空字符串
		// username := ctx.DefaultPostForm("username", "somebody")
		// password := ctx.DefaultPostForm("xxx", "****")

		username, ok := ctx.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := ctx.GetPostForm("password")
		if !ok {
			password = "***"
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run()
}
