package middleware

import (
	"gin-web/common"
	"gin-web/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			// 不符合格式
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足！"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足！"})
			c.Abort()
			return
		}
		// 验证通过，获取claims中userId
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		// 如果用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "权限不足！"})
			c.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}
