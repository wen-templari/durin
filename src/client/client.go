package client

import (
	"durin/src/model"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id      string
	conn    *websocket.Conn
	send    chan model.Message
	manager *ClientManager
}

// var ConnectionMap map[string]*Client = make(map[string]*Client)

func (c *Client) write() {
	defer c.close()
	for {
		message := <-c.send
		err := c.conn.WriteJSON(message)
		if err != nil {
			break
		}
	}
}

func (c *Client) read() {
	defer c.close()
	for {
		message := model.Message{}
		_, r, err := c.conn.NextReader()
		if err != nil {
			log.Println(err)
			break
		}
		err = json.NewDecoder(r).Decode(&message)
		if err != nil {
			log.Println(err)
		} else {
			c.manager.Send <- message
		}
	}
}
func (c *Client) close() {
	c.manager.Unregister <- c.Id
	c.conn.Close()
}

func SaveClient(id string, conn *websocket.Conn) int {
	_, ok := Manager.Clients[id]
	if ok {
		log.Println("Client already exist")

	} else {
		log.Println("Client added")
		c := &Client{id, conn, make(chan model.Message), &Manager}
		Manager.Register <- c
		go c.write()
		go c.read()
	}
	return 0
}
