package client

import (
	"durin/model"
	"durin/util"
	"encoding/json"
	"log"
)

type ClientManager struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan string
	Send       chan model.Message
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client),
	Register:   make(chan *Client),
	Unregister: make(chan string),
	Send:       make(chan model.Message),
}

func (manager *ClientManager) Run() {
	defer log.Println("ClientManager stopped")
	for {
		select {
		case client := <-manager.Register:
			// manager.Clients[client] = true
			if _, ok := manager.Clients[client.Id]; !ok {
				manager.Clients[client.Id] = client
				log.Printf("Added new client %s\n", client.Id)
			} else {
				manager.Clients[client.Id].conn = client.conn
				log.Println("Client already exist, updating conn")
			}
		case id := <-manager.Unregister:
			if _, ok := manager.Clients[id]; ok {
				log.Printf("Removed client %s\n", id)
				close(manager.Clients[id].send)
				delete(manager.Clients, id)
			}
		case message := <-manager.Send:
			if manager.Clients[message.To] != nil {
				log.Println("sending")
				manager.Clients[message.To].send <- message
			} else {
				conn := util.Pool.Get()
				log.Printf("Client %s not found\n", message)
				messageStr, _ := json.Marshal(message)
				conn.Do("RPUSH", "stashed:"+message.To, messageStr)
				conn.Close()
			}
		}
	}
}
