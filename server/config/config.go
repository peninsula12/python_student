package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/python_student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = *db
}
