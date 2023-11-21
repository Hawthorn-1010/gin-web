package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Telephone string
	Password  string
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()

	r.POST("/api/auth/register", func(c *gin.Context) {
		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		// 数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "手机号格式不正确！"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "密码格式不正确！"})
			return
		}
		if len(name) == 0 {
			// 如果用户名字为空，随机生成一个10位数string
			name = randomString(10)
		}
		// 判断手机号是否存在
		if isUserExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在！"})
			return
		}
		db.Create(&User{Name: name, Telephone: telephone, Password: password})

		// 返回结果
		c.JSON(200, gin.H{"msg": "注册成功"})
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func randomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	// 10位数组
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(192.168.255.3:3306)/books"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 自动创建数据表
	db.AutoMigrate(&User{})
	return db
}

func isUserExist(db *gorm.DB, telephone string) bool {
	var user User
	db.First(&user, "Telephone = ?", telephone) // 查找
	if user.ID != 0 {
		return true
	}
	return false
}

//// 迁移 schema
//db.AutoMigrate(&User{})
//
//// Create
//db.Create(&User{Name: name, Telephone: telephone, Password: password})
//
//// Read
//var user User
//db.First(&user, 1)               // 根据整型主键查找
//db.First(&user, "Name = ?", "h") // 查找 code 字段值为 D42 的记录
//
//// Update - 将 product 的 price 更新为 200
//db.Model(&user).Update("Price", 200)
//// Update - 更新多个字段
//db.Model(&user).Updates(User{Telephone: "12345678901"}) // 仅更新非零值字段
////db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
//
//// Delete - 删除 product
//db.Delete(&user, 1)
