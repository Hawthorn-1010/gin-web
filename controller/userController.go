package controller

import (
	"01-quickstart/common"
	"01-quickstart/dto"
	"01-quickstart/model"
	"01-quickstart/response"
	"01-quickstart/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)
	// 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号格式不正确")
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
	newUser := model.User{Name: name, Telephone: telephone, Password: password}
	db.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常！"})
		log.Printf("token generate error %v", err)
		return
	}

	// 返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})

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
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常！"})
		log.Printf("token generate error %v", err)
		return
	}

	// 返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	// 返回结果
	response.Response(c, 200, 200, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}, "获取用户信息成功")

}

func isUserExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.First(&user, "Telephone = ?", telephone) // 查找
	if user.ID != 0 {
		return true
	}
	return false
}
