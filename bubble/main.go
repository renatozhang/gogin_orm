package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDb, err := DB.DB()
	if err != nil {
		return
	}
	return sqlDb.Ping()
}

func main() {
	// 创建数据库
	// sql: create database bubble
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	router := gin.Default()
	// 加载静态文件
	router.Static("/static/", "static")
	// 加载模板文件
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := router.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			// 前端页面填写待办事项,点击提交,挥发请求到这里
			// 1.从请求中把数据拿出来
			var todo Todo
			ctx.BindJSON(&todo)
			// 2.存储数据库
			if err = DB.Create(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
				// ctx.JSON(http.StatusOK, gin.H{
				// 	"code": 2000,
				// 	"msg":  "success",
				// 	"data": todo,
				// })
			}
			// 3.返回相应

		})
		// 查看
		v1Group.GET("/todo", func(ctx *gin.Context) {
			// 查询todo这个表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, _ := ctx.Params.Get("id")
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			ctx.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}

		})
		// 删除
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"id": "deleted"})
			}

		})

	}

	router.Run(":9000")
}
