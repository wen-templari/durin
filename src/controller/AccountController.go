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

	i, err := conn.Do("EXISTS", account.Id)
	if err != nil {
		log.Println(err)
	}
	if i.(int64) == 0 {
		c.JSON(200, util.NewReturnObject(400, "id is not exist", nil))
		return
	}

	n, err := redis.Values(conn.Do("hgetall", account.Id))
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
	conn.Do("set", "token:"+account.Id, token)
	conn.Do("expire", "token:"+account.Id, 43200)
	c.SetCookie("token", token, 3600, "/", "", false, false)

	//TODO return stashed message

	c.JSON(200, util.NewReturnObject(200, "login success", dbResult))
}

func Register(c *gin.Context) {
	var account model.Account
	err := c.ShouldBind(&account)

	if err != nil {
		log.Println(err)
		return
	}

	if account.Name == "" || account.Password == "" {
		c.JSON(200, util.NewReturnObject(400, "name or password is empty", nil))
		return
	}

	conn := util.Pool.Get()
	defer conn.Close()

	// self increase id
	id, err := conn.Do("incr", "idCounter")
	if err != nil {
		log.Println(err)
	}

	conn.Do("hmset", id, "id", id, "name", account.Name, "password", account.Password)

	c.JSON(200, util.NewReturnObject(200, "register success", nil))
}

func Logout(c *gin.Context) {
}

func Search(c *gin.Context) {
	id, _ := c.GetQuery("id")
	conn := util.Pool.Get()
	defer conn.Close()

	n, err := redis.Values(conn.Do("hgetall", id))
	if err != nil {
		log.Println(err)
		c.JSON(200, util.NewReturnObject(500, "server error", nil))
		return
	}

	var dbResult model.AccountDTO
	err = redis.ScanStruct(n, &dbResult)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, util.NewReturnObject(200, "success", dbResult))

}
