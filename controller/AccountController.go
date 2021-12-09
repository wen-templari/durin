package controller

import (
	"durin/client"
	"log"

	"github.com/gin-gonic/gin"
)

type Account struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var account Account = Account{}
	if c.ShouldBind(&account) == nil {
		log.Printf(account.Username)
	}
	newClient := &client.Client{}
	newClient.Id = account.Id
	client.ConnectionMap[account.Id] = *newClient
}

func Register(c *gin.Context) {
	var account Account = Account{}
	if c.ShouldBind(&account) == nil {
		log.Printf(account.Username)
	}
}
