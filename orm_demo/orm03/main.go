package main

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 定义模型
type User struct {
	ID int64
	// Name *string `gorm:"default:'小王子'"`
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	// u1 := User{Name: "qimi", Age: 18}
	// tx := db.Create(&u1)
	// fmt.Println(tx.RowsAffected)

	// u2 := User{Name: "", Age: 19} // default 默认值
	// db.Create(&u2)

	// u3 := User{Name: new(string), Age: 39}
	// db.Create(&u3)

	u4 := User{Name: sql.NullString{String: "", Valid: true}, Age: 10}
	db.Create(&u4) // 此时数据库中该条记录name字段的值就是''

	// 查询
	// var user User
	// db.First(&user)
	// fmt.Printf("user:%#v", user)

}
