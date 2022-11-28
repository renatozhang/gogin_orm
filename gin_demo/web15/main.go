package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20 //8 MiB
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(ctx *gin.Context) {
		// 从请求中读取文件
		f, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		// 将读取到的文件保存到本地（服务器端）
		// dst := fmt.Sprintf("./%s", f.Filename)
		dst := path.Join("./", f.Filename)
		err = ctx.SaveUploadedFile(f, dst)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/uploads", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Printf(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			// 上传文件到指定目录
			ctx.SaveUploadedFile(file, dst)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run()
}
