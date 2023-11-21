package controller

import (
	"01-quickstart/common"
	"01-quickstart/model"
	"01-quickstart/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func Register(c *gin.Context) {
	db := common.GetDB()
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
		name = util.RandomString(10)
	}
	// 判断手机号是否存在
	if isUserExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在！"})
		return
	}
	db.Create(&model.User{Name: name, Telephone: telephone, Password: password})

	// 返回结果
	c.JSON(200, gin.H{"msg": "注册成功"})
}

func Login(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	var user model.User
	db.First(&user, "Telephone = ?", telephone) // 查找

	// 判断手机号是否存在
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在，请注册！"})
		return
	}
	// 验证密码是否正确
	if strings.Compare(user.Password, password) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码错误！"})
		return
	}
	// 发放token
	token := "11"
	// 返回结果
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})
}

func isUserExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.First(&user, "Telephone = ?", telephone) // 查找
	if user.ID != 0 {
		return true
	}
	return false
}
