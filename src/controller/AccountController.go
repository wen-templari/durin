package controller

import (
	"durin/src/model"
	"durin/src/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Login(c *gin.Context) {
	var account model.Account = model.Account{}
	c.ShouldBind(&account)
	if account.Id == "" || account.Password == "" {
		c.JSON(200, util.NewReturnObject(400, "id or password is empty", nil))
		return
	}

	conn := util.Pool.Get()
	defer conn.Close()

	i, err := conn.Do("EXISTS", "user:"+account.Id)
	if err != nil {
		log.Println(err)
	}
	if i.(int64) == 0 {
		c.JSON(200, util.NewReturnObject(400, "id is not exist", nil))
		return
	}

	n, err := redis.Values(conn.Do("hgetall", "user:"+account.Id))
	if err != nil {
		log.Println(err)
		c.JSON(200, util.NewReturnObject(500, "server error", nil))
		return
	}

	var dbResult model.Account
	err = redis.ScanStruct(n, &dbResult)
	if err != nil {
		log.Println(err)
	}
	if dbResult.Password != account.Password {
		c.JSON(200, util.NewReturnObject(400, "wrong password", nil))
		return
	}

	token := util.GenerateToken()
	former, _ := c.Cookie("token")
	log.Println(former)
	conn.Do("set", "user:token:"+account.Id, token)
	conn.Do("expire", "user:token:"+account.Id, 43200)
	c.SetCookie("token", token, 3600, "/", "", false, false)

	//TODO return stashed message

	c.JSON(200, util.NewReturnObject(200, "login success", dbResult))
}

func Register(c *gin.Context) {
	var account model.Account
	if c.ShouldBind(&account) == nil {
		log.Printf(account.Name)
	}
}
