package client

import (
	"durin/model"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	mu      sync.Mutex
	Id      string
	conn    *websocket.Conn
	send    chan model.Message
	manager *ClientManager

	resetHeartbeatTimer bool
}

func (c *Client) write() {
	defer c.close()
	for {
		message := <-c.send
		log.Println(message)
		if len(message.From) == 0 {
			break
		}
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
			if message.To == "heartbeat" {
				c.resetHeartbeatTimer = true
			} else {
				c.manager.Send <- message
			}
		}
	}
}

func (c *Client) heartbeatTimer() {
	fragmentCount := 600
	unregisterTime := 5000
	for i := 0; i < fragmentCount; i++ {
		c.mu.Lock()
		if c.resetHeartbeatTimer {
			c.resetHeartbeatTimer = false
			i = 0
			c.mu.Unlock()
			continue
		}
		c.mu.Unlock()
		time.Sleep(time.Duration(unregisterTime/fragmentCount) * time.Millisecond)
	}
	c.close()
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
		// c := &Client{id, conn, make(chan model.Message), &Manager}
		c := &Client{
			Id:      id,
			conn:    conn,
			send:    make(chan model.Message),
			manager: &Manager,

			resetHeartbeatTimer: false,
		}
		Manager.Register <- c
		go c.write()
		go c.read()
		go c.heartbeatTimer()
	}
	return 0
}
