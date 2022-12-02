package dao

import (
	"fmt"

	"github.com/renatozhang/gogin_orm/bubblev1/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	fmt.Println(dsn)
	// dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
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

func Close() {
	sqlDb, err := DB.DB()
	if err != nil {
		return
	}
	sqlDb.Close()
}
