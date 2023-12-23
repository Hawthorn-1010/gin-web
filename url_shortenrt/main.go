package main

import (
	"github.com/gin-gonic/gin"
	"urlshortener/service"
)

func main() {
	service.InitClient()
	r := gin.Default()
	r = CollectRoute(r)
	// Listen and Server in 0.0.0.0:8080
	//r.GET("/hello", func(context *gin.Context) {
	//	context.Writer.Write([]byte("hello!!!"))
	//})
	r.Run(":8081")
}
