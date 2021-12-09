package main

import (
	"github.com/gin-gonic/gin"

	"durin/controller"
)

func main() {
	bindAddress := "localhost:8080"
	r := gin.Default()
	r.GET("/test", controller.Ping)
	r.POST("/login", controller.Login)
	r.Run(bindAddress)
}
