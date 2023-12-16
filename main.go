package main

import (
	"01-quickstart/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}

//// 迁移 schema
//db.AutoMigrate(&User{})
//
//// Create
//db.Create(&User{Name: name, Telephone: telephone, Password: password})
//
//// Read
//var model User
//db.First(&model, 1)               // 根据整型主键查找
//db.First(&model, "Name = ?", "h") // 查找 code 字段值为 D42 的记录
//
//// Update - 将 product 的 price 更新为 200
//db.Model(&model).Update("Price", 200)
//// Update - 更新多个字段
//db.Model(&model).Updates(User{Telephone: "12345678901"}) // 仅更新非零值字段
////db.Model(&model).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
//
//// Delete - 删除 product
//db.Delete(&model, 1)
