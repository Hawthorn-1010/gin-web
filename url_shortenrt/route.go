package main

import (
	"github.com/gin-gonic/gin"
	"urlshortener/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//r.Use(middleware.CORSMiddleware())
	r.POST("/create_url", controller.SetUrl)
	r.GET("/:short_url", controller.GetUrl)
	//r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
