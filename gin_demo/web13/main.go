package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取请求的path(URL)参数
// 注意URL的匹配不要冲突
func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(ctx *gin.Context) {
		// 获取路径参数
		name := ctx.Param("name")
		age := ctx.Param("age")
		ctx.JSON(http.StatusOK, gin.H{
			"username": name,
			"age":      age,
		})
	})

	r.GET("/blog/:year/:month", func(ctx *gin.Context) {
		year := ctx.Param("year")
		month := ctx.Param("month")
		ctx.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run()
}
