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
	r.POST("/user", controller.Register)
	r.GET("/user", controller.Search)
	r.POST("/user/login", controller.Login)
	r.POST("/user/:id/avatar", controller.SetAvatar)
	r.GET("/message", controller.Message)
	r.POST("/file", controller.File)
	r.Run(bindAddress)
}
