package main

import "github.com/gin-gonic/gin"

func main() {
	//常见一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// fanhui Json 格式的数据
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})

	r.Run()
}
