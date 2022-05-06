package controller

import (
	"durin/model"
	"durin/util"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Login(c *gin.Context) {
	var account model.Account = model.Account{}
	c.ShouldBind(&account)
	if account.Id == "" || account.Password == "" {
		c.JSON(200, util.NewReturnObject(1201, "id or password is empty", nil))
		return
	}

	conn := util.Pool.Get()
	defer conn.Close()

	i, err := conn.Do("EXISTS", account.Id)
	if err != nil {
		log.Println(err)
	}
	if i == nil || i.(int64) == 0 {
		c.JSON(200, util.NewReturnObject(1401, "id is not exist", nil))
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
		c.JSON(200, util.NewReturnObject(1301, "wrong password", nil))
		return
	}

	token := util.GenerateToken()
	former, _ := c.Cookie("token")
	log.Println(former)
	conn.Do("set", "token:"+account.Id, token)
	conn.Do("expire", "token:"+account.Id, 43200)
	c.SetCookie("token", token, 3600, "/", "", false, false)

	offlineMessageString, _ := redis.Strings(conn.Do("lrange", "stashed:"+account.Id, 0, -1))
	if len(offlineMessageString) > 0 {
		conn.Do("del", "stashed:"+account.Id)
	}
	offlineMessage := []model.Message{}
	for _, v := range offlineMessageString {
		temp := model.Message{}
		err := json.Unmarshal([]byte(v), &temp)
		if err != nil {
			log.Println(err)
		}
		offlineMessage = append(offlineMessage, temp)
	}

	fmt.Println(offlineMessage)

	accountDTO := model.AccountDTO{
		Id:             dbResult.Id,
		Name:           dbResult.Name,
		Avatar:         dbResult.Avatar,
		Token:          token,
		OfflineMessage: offlineMessage,
	}

	c.JSON(200, util.NewReturnObject(0, "login success", accountDTO))
}

func Register(c *gin.Context) {
	var account model.Account
	err := c.ShouldBind(&account)

	if err != nil {
		log.Println(err)
		return
	}

	if account.Name == "" || account.Password == "" {
		c.JSON(200, util.NewReturnObject(1201, "name or password is empty", nil))
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
	log.Println(id.(int64))
	strId := strconv.FormatInt(id.(int64), 10)
	accountDTO := model.AccountDTO{
		Id:   strId,
		Name: account.Name,
	}

	c.JSON(200, util.NewReturnObject(0, "register success", accountDTO))
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

	c.JSON(200, util.NewReturnObject(0, "success", dbResult))

}

func SetAvatar(c *gin.Context) {
	id := c.Param("id")
	file, err := c.FormFile("upload")
	if err != nil {
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		log.Println(err)
	}

	filePath := util.UploadImg(openedFile)

	conn := util.Pool.Get()
	defer conn.Close()

	conn.Do("hset", id, "avatar", filePath)
	fileDTO := model.FileDTO{
		Path: filePath,
	}
	c.JSON(200, util.NewReturnObject(0, "Success", fileDTO))
}
