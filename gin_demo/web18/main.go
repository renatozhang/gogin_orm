package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件

func indexHandler(ctx *gin.Context) {
	fmt.Println("index")
	name, ok := ctx.Get("name") //从上下文中取值
	if !ok {
		name = "匿名用户"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1:统计请求处理函数的耗时

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	// 记时
	start := time.Now()
	// go funcxx(c.Copy()) //在funcxx中只能使用c的拷贝，不能使用C
	c.Next() // 调用后续处理函数
	// c.Abort() //组织调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "q1mi") // 在上下文c中设置值
	c.Next()
	// c.Abort() //组织后续的处理函数
	// return
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if doCheck {
			// 是否登录的判断
			// if 是登录用户
			// c.Next()
			//else
			// c.Abort()
		} else {
			ctx.Next()
		}

	}
}

func main() {
	// router := gin.Default() //默认使用了Logger(), Recovery()
	router := gin.New()
	router.Use(m1, m2) // 全局注册m1,m2
	router.GET("/index", indexHandler)
	router.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	// 路由组注册中间件方法1：
	xxGroup := router.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "/xx/index",
			})
		})
	}
	// 路由组注册中间件方法2：
	xx2Group := router.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "/xx/index",
			})
		})
	}
	router.Run()
}
