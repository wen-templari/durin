package controller

import (
	"durin/src/client"
	"durin/src/util"
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

	i, err := conn.Do("EXISTS", "user:"+id)
	if err != nil {
		log.Println(err)
	}
	if i.(int64) == 0 {
		c.JSON(200, util.NewReturnObject(400, "id is not exist", nil))
		return
	}
	tokenGet, err := redis.String(conn.Do("get", "user:token:"+id))
	if err != nil {
		log.Println(err)
		c.JSON(200, util.NewReturnObject(500, "server error", nil))
		return
	}

	// if n == c.Cookie("token") {
	// 	c.JSON(200, util.NewReturnObject(400, "token is not exist", nil))
	// 	return
	// }

	tokenSend, _ := c.Cookie("token")
	if tokenSend == "" {
		c.JSON(200, util.NewReturnObject(400, "token is not exist", nil))
		return
	}

	if tokenSend != tokenGet {
		c.JSON(200, util.NewReturnObject(400, "wrong token", nil))
		return
	}

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// log.Println(wsConn)
	client.SaveClient(id, wsConn)
}
