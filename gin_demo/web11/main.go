package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// GET请求 URL ?后面的是querystring参数
	// key=value格式，多个key-value用 & 连接
	// eq:  /web/query=小王子&age=18
	r := gin.Default()
	r.GET("/web", func(ctx *gin.Context) {
		// 获取浏览器请求携带的query string 参数
		name := ctx.Query("query") //通过Query获取请求中携带的querystring参数
		// name := ctx.DefaultQuery("query", "somebody") //取不到使用指定的默认值
		// name, ok := ctx.GetQuery("query") // 取到返回(值, true)，取不到返回("", false)
		// if !ok {
		// 	// 取不到
		// 	name = "somebody"
		// }
		age := ctx.Query("age")

		ctx.JSON(http.StatusOK, gin.H{
			"query": name,
			"age":   age,
		})

	})

	r.Run()
}
