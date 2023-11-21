package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// middleware，不管访问哪个请求都会执行到
	r.Use(func(context *gin.Context) {
		nowtime := time.Now()

		logger.Info("Incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(nowtime)),
		)
		// 继续执行
		context.Next()
	}, func(c *gin.Context) {
		c.Set("Id", rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {

		h := gin.H{
			"message": "pong",
		}

		if rid, exist := c.Get("Id"); exist {
			h["Id"] = rid
		}

		c.JSON(200, h)
	})
	r.Run()
}
