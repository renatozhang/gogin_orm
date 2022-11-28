package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	// Any: 请求方法的大集合/大杂烩
	r.Any("/user", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case "GET":
			ctx.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			ctx.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})
	})

	// NoRoute 设置404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "404 not found"})
	})

	// route group
	// 视频的首页和详情页
	// r.GET("/video/index", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	// })
	// r.GET("/video/xx", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
	// })
	// r.GET("/video/oo", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
	// })

	// 路由组的组 多用于区分不同的业务线或API版本
	// 把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})

		videoGroup.GET("/xx", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})

		videoGroup.GET("/oo", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}

	//商城的首页和详情页
	r.GET("/shop/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})

	r.Run()
}
