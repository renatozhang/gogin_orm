package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gogin_orm/bubblev1/controller"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()
	// 加载静态文件
	router.Static("/static/", "static")
	// 加载模板文件
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandller)

	// v1
	v1Group := router.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.GET("/todo/:id", controller.GetATodo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return router
}
