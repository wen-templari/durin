package controller

import (
	"durin/client"
	"durin/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Message(c *gin.Context) {
	id, _ := c.GetQuery("id")

	conn := util.Pool.Get()
	defer conn.Close()

	log.Println(id)

	i, err := conn.Do("EXISTS", id)
	if err != nil {
		log.Println(err)
	}
	if i.(int64) == 0 {
		c.JSON(200, util.NewReturnObject(1401, "id is not exist", nil))
		return
	}
	tokenGet, err := redis.String(conn.Do("get", "token:"+id))
	if err != nil {
		log.Println(err)
		c.JSON(200, util.NewReturnObject(500, "server error", nil))
		return
	}

	tokenSend, _ := c.GetQuery("token")
	if tokenSend == "" {
		c.JSON(200, util.NewReturnObject(1421, "token is not exist", nil))
		return
	}

	if tokenSend != tokenGet {
		c.JSON(200, util.NewReturnObject(1321, "wrong token", nil))
		return
	}

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client.SaveClient(id, wsConn)
}
