package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Query(ctx *gin.Context)
}
