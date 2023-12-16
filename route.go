package main

import (
	"gin-web/controller"
	"gin-web/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/category")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Query)

	postRoutes := r.Group("/post")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Query)
	postRoutes.POST("/page/list", postController.PageList)

	return r
}
