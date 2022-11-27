package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(ctx *gin.Context) {
		// 方法1: 使用map
		// data := map[string]interface{}{
		// 	"name":    "小王子",
		// 	"message": "hello world",
		// 	"age":     18,
		// }

		data := gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		}
		ctx.JSON(http.StatusOK, data)
	})
	// 方法2：结构体 灵活使用tag来对结构体字段做定制化
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}

	r.GET("/anthor_json", func(ctx *gin.Context) {
		data := msg{
			Name:    "小王子",
			Message: "hello golang",
			Age:     18,
		}
		ctx.JSON(http.StatusOK, data) // json的序列化
	})

	r.Run()
}
