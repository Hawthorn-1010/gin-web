package common

import (
	"01-quickstart/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(192.168.255.3:3306)/books?parseTime=true"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 自动创建数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
