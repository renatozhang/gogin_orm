package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gogin_orm/bubblev1/model"
)

/*
	url--> controller  -->logic --> model
	请求来了 -->控制器 --> 业务逻辑 --> 模型层的增删改查
*/

func IndexHandller(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(ctx *gin.Context) {
	// 前端页面填写待办事项,点击提交,挥发请求到这里
	// 1.从请求中把数据拿出来
	var todo model.Todo
	ctx.BindJSON(&todo)
	// 2.存储数据库
	err := model.CreateATodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(ctx *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := model.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func GetATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := model.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, todo)
	}

}

func UpdateATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	todo, err := model.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.BindJSON(&todo)
	if err = model.UpdateATodo(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := model.DeleteATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id": "deleted"})
	}
}
