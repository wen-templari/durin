package main

import (
	"durin/src/client"
	"durin/src/controller"
	"durin/src/util"

	"github.com/gin-gonic/gin"
)

func main() {

	util.InitPool()
	go client.Manager.Run()
	bindAddress := "localhost:8080"
	r := gin.Default()
	r.POST("/login", controller.Login)
	r.GET("/message", controller.Message)
	r.Run(bindAddress)
}
