package client

import (
	"durin/src/model"
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
				log.Println("Client already exist")
			}
		case id := <-manager.Unregister:
			if _, ok := manager.Clients[id]; ok {
				close(manager.Clients[id].send)
				delete(manager.Clients, id)
			}
		case message := <-manager.Send:
			// err := ConnectionMap[message.To].Conn.WriteJSON(message)
			if manager.Clients[message.To] != nil {
				manager.Clients[message.To].send <- message
			}
		}
	}
}
