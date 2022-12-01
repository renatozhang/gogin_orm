package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // ID CreatedAt  UpdatedAt  DeletedAt
	Name       string
	Age        int64
}

func main() {
	// 1.连接MySQL数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 2.把模型与数据库中的表对应起来
	// db.AutoMigrate(&User{})

	// 3.创建
	// u1 := User{Name: "q1mi", Age: 18}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu", Age: 20}
	// db.Create(&u2)

	// 查询
	// var user User
	// 根据主键查询第一条数据
	// user := new(User) // new和make的区别
	// db.Debug().First(user)
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	// fmt.Printf("user:%#v\n", user)

	// var users []User
	// db.Debug().Find(&users) // SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	// fmt.Printf("users:%#v\n", users)

	//Where
	// db.Debug().Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM `users` WHERE name = 'jinzhu' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1

	// db.Debug().Where("name=?", "jinzhu").Find(&users)
	// SELECT * FROM `users` WHERE name='jinzhu' AND `users`.`deleted_at` IS NULL

	// db.Debug().Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM `users` WHERE name <> 'jinzhu' AND `users`.`deleted_at` IS NULL

	// db.Debug().Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM `users` WHERE name IN ('jinzhu','jinzhu 2') AND `users`.`deleted_at` IS NULL

	// db.Debug().Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM `users` WHERE name LIKE '%jin%' AND `users`.`deleted_at` IS NULL

	// db.Debug().Where("name = ? AND age>=?", "jinzhu", 22).Find(&users)
	// SELECT * FROM `users` WHERE (name = 'jinzhu' AND age>=22) AND `users`.`deleted_at` IS NULL

	// lastweek := time.Now().AddDate(0, 0, -7)
	// db.Debug().Where("updated_at > ?", lastweek).Find(&users)
	// SELECT * FROM `users` WHERE updated_at > '2022-11-24 11:59:41.509' AND `users`.`deleted_at` IS NULL
	// fmt.Printf("users:%#v\n", users)

	// BETWEEN
	// db.Debug().Where("created_at BETWEEN ? AND ?", lastweek, time.Now()).Find(&users)
	// SELECT * FROM `users` WHERE (created_at BETWEEN '2022-11-24 11:59:13.479' AND '2022-12-01 11:59:13.479') AND `users`.`deleted_at` IS NULL

	// Struct & Map
	// struct
	// var user User
	// db.Debug().Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu' AND `users`.`age` = 20 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1

	// map
	// db.Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM `users` WHERE `age` = 20 AND `name` = 'jinzhu' AND `users`.`deleted_at` IS NULLs

	//主键的切片
	// db.Debug().Where([]int64{20, 12, 22}).Find(&users)
	// SELECT * FROM `users` WHERE `users`.`id` IN (20,12,22) AND `users`.`deleted_at` IS NULL

	// 字段值为0,'',false或其他零值,将不会用于构建查询条件
	// db.Debug().Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu'

	// FirstOrCreate
	// var user User
	// db.Debug().FirstOrCreate(&user, User{Name: "小王子"})

	// Attrs && assign
	// var user User
	// db.Debug().Attrs(User{Age: 99}).FirstOrCreate(&user, User{Name: "小王子"})
	// db.Debug().Assign(User{Age: 99}).FirstOrCreate(&user, User{Name: "小王子"})

}
