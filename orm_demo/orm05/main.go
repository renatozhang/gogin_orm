package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // ID CreatedAt  UpdatedAt  DeletedAt
	Name       string
	Age        int64
	Active     bool
}

func main() {
	// 2.连接MySQL数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// // 3.把模型与数据库中的表对应起来
	// db.AutoMigrate(&User{})

	// // 4.创建
	// u1 := User{Name: "q1mi", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu", Age: 20, Active: false}
	// db.Create(&u2)

	// // 5.查询
	var user User
	db.First(&user)

	// // 6.更新
	user.Name = "七米"
	user.Age = 99
	// db.Debug().Save(&user) //默认会修改所有字段
	// UPDATE `users` SET `created_at`='2022-12-01 13:56:25.448',`updated_at`='2022-12-01 14:00:31.345',`deleted_at`=NULL,`name`='七米',`age`=99,`active`=true

	// 更新单个属性,如果他有变化
	// db.Debug().Model(&user).Update("name", "小王子")
	// PDATE `users` SET `name`='小王子',`updated_at`='2022-12-01 14:02:47.066' WHERE `users`.`deleted_at` IS NULL AND `id` = 1

	// m1 := map[string]interface{}{"name": "zhangzeng", "age": 18, "active": true}
	// db.Debug().Model(&user).Updates(m1) //m1所有列都会更新
	// UPDATE `users` SET `active`=true,`age`=18,`name`='zhangzeng',`updated_at`='2022-12-01 14:10:15.557' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	// db.Debug().Model(&user).Select("age").Updates(m1) //只更新age字段
	// UPDATE `users` SET `age`=18,`updated_at`='2022-12-01 14:07:57.933' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	// db.Debug().Model(&user).Omit("active").Updates(m1) //排除m1中的active字段
	// UPDATE `users` SET `age`=18,`name`='zhangzeng',`updated_at`='2022-12-01 14:09:07.259' WHERE `users`.`deleted_at` IS NULL AND `id` = 1

	// db.Debug().Model(&user).UpdateColumn("age", 30)
	// UPDATE `users` SET `age`=30 WHERE `users`.`deleted_at` IS NULL AND `id` = 1

	// rowsNum := db.Debug().Model(User{}).Updates(User{Name: "golang", Age: 18}).RowsAffected
	// fmt.Println(rowsNum)

	// 让users表中所有的用户的年龄在原来的基础上+2
	// db.Debug().Where("age > ?", 0).Model(&User{}).Update("age", gorm.Expr("age  + ?", 2))
}
