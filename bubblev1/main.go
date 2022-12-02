package main

import (
	"fmt"
	"log"

	"github.com/renatozhang/gogin_orm/bubblev1/dao"
	"github.com/renatozhang/gogin_orm/bubblev1/model"
	"github.com/renatozhang/gogin_orm/bubblev1/routers"
	"github.com/renatozhang/gogin_orm/bubblev1/setting"
)

func main() {
	if err := setting.Init("./conf/config.ini"); err != nil {
		log.Printf("%#v", setting.Conf.MySQLConfig)
		panic(err)
	}

	// 创建数据库
	// sql: create database bubble
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&model.Todo{})

	//注册路由
	router := routers.SetupRouters()

	if err := router.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v", err)
	}
}
