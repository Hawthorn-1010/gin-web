package common

import (
	"fmt"
	"gin-web/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	loc := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		url.QueryEscape(loc))
	db, err := gorm.Open(driverName, dsn)
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
