package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	r.Run()
}
