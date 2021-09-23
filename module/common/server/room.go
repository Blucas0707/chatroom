package server

import (
	"fmt"
	"log"
)

const welcomeMessage = "%s joined the room"

type Room struct {
	name       string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

func NewRoom(name string) *Room {
	return &Room{
		name:       name,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
}

func (room *Room) RunRoom() {
	fmt.Println("room.go RunRoom executing")
	for {
		select {
		case client := <-room.register:
			room.registerClientInRoom(client)
		case client := <-room.unregister:
			room.unregisterClientInRoom(client)
		case message := <-room.broadcast:
			fmt.Println("RunRoom :", message)
			room.broadcastToClientsInRoom(message.encode())
		}
	}
}

func (room *Room) registerClientInRoom(client *Client) {
	room.notifyClientJoined(client)
	room.clients[client] = true
}

func (room *Room) unregisterClientInRoom(client *Client) {
	if _, ok := room.clients[client]; !ok {
		fmt.Printf("Client %s left the room", client.Name)
		delete(room.clients, client)
	}
}

func (room *Room) broadcastToClientsInRoom(messageEncode []byte) {
	for client := range room.clients {
		client.send <- messageEncode
	}
}

func (room *Room) notifyClientJoined(client *Client) {
	message := &Message{
		Action:  SendMessageAction,
		Target:  room.name,
		Message: fmt.Sprintf(welcomeMessage, client.GetName()),
	}
	log.Println("notify message: ", message)
	room.broadcastToClientsInRoom(message.encode())
}

func (room *Room) GetName() string {
	return room.name
}
