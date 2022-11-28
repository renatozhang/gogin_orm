package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		// 跳转到别的网站
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/a", func(ctx *gin.Context) {
		//跳转到/b对应的路由函数
		// 指定重定向的URL
		ctx.Request.URL.Path = "/b"
		r.HandleContext(ctx)

	})

	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "b",
		})

	})
	r.Run()
}
