package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     int64
	Name   string
	Gender string
	Hobby  string
}

// gorm demo1
func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Printf("connection database failed, err:%v\n", err)
		return
	}
	// 创建数据行
	// db.AutoMigrate(&UserInfo{})
	// u1 := UserInfo{ID: 1, Name: "七米", Gender: "男", Hobby: "蛙泳"}
	// db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //查询表中第一条数据保存到u中
	fmt.Printf("u:%#v", u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)

}
