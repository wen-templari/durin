package client

import "github.com/gorilla/websocket"

type Client struct {
	Id   string
	Conn *websocket.Conn
	Send chan []byte
}

var ConnectionMap map[string]Client = make(map[string]Client)

func (c *Client) write() {
	for {
		message := <-c.Send
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

func (c *Client) read() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		c.Send <- message
	}
}

func AddConnection(id string, conn *websocket.Conn) int {
	// c := &Client{id, conn, make(chan []byte)}
	c, ok := ConnectionMap[id]
	if !ok {
		return -1
	}
	c.Conn = conn
	c.Send = make(chan []byte)
	go c.write()
	go c.read()
	return 0
}
