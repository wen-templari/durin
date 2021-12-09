package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ping(c *gin.Context) {
	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	return
	// }
	// addConnection("1", conn)
}
