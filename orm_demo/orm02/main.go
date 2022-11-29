package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model                 //内嵌gorm.Model
	Name         string        `gorm:"size:255"`
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);uniqueIndex"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type Animal struct {
	AnimalID int64 `gorm:"primaryKey"`
	Name     string
	Age      int64
}

// func (Animal) TableName() string {
// 	return "qimi"
// }

func main() {
	// &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}, 表明不加s
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("connection database failed, err:%v\n", err)
		return
	}

	// db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
}
