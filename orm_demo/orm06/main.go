package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model // ID CreatedAt  UpdatedAt  DeletedAt
// 	Name       string
// 	Age        int64
// 	Active     bool
// }

type User struct {
	ID     int64
	Name   string
	Age    int64
	Active bool
}

func main() {
	// 2.连接MySQL数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 	// 3.把模型与数据库中的表对应起来
	// db.AutoMigrate(&User{})

	// // 4.创建
	// u1 := User{Name: "q1mi2", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu2", Age: 20, Active: false}
	// db.Create(&u2)

	// 删除
	// var u User
	// u.ID = 1
	// db.Debug().Delete(&u)
	// // UPDATE `users` SET `deleted_at`='2022-12-01 14:41:53.137' WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL

	// var u User
	// u.Name = "jinzhu2"
	// db.Debug().Delete(&u)
	// UPDATE `users` SET `deleted_at`='2022-12-01 14:43:07.414' WHERE `users`.`deleted_at` IS NULL

	db.Where("name=?", "jinzhu2").Delete(&User{})
	// UPDATE `users` SET `deleted_at`='2022-12-01 14:44:10.052' WHERE name='jinzhu2' AND `users`.`deleted_at` IS NULL

	// db.Debug().Delete(&User{}, "age=?", 18)
	// db.Debug().Where("name=?", "jinzhu2").Delete(&User{})

	// 物理删除
	// var users []User
	// db.Debug().Unscoped().Where("name=?", "jinzhu2").Find(&users)
	// SELECT * FROM `users` WHERE name='jinzhu2'

	// db.Debug().Unscoped().Where("name=?", "jinzhu2").Delete(&User{})
	// DELETE FROM `users` WHERE name='jinzhu2'

}
