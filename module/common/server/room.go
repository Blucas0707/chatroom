package server

import (
	"fmt"
	"log"
)

const welcomeMessage = "%s joined the room"
const leaveMessage = "%s left the room"

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
	if _, ok := room.clients[client]; !ok {
		room.clients[client] = true
		room.notifyClientJoined(client)
	}
}

func (room *Room) unregisterClientInRoom(client *Client) {
	if _, ok := room.clients[client]; ok {
		fmt.Printf("Client %s left the room\n", client.Name)
		delete(room.clients, client)
		room.notifyClientLeft(client)
	}
}

func (room *Room) broadcastToClientsInRoom(messageEncode []byte) {
	for client := range room.clients {
		client.send <- messageEncode
	}
}

func (room *Room) notifyClientJoined(client *Client) {
	message := &Message{
		Action:  JoinRoomAction,
		Target:  room.name,
		Message: fmt.Sprintf(welcomeMessage, client.GetName()),
		User:    room.ListRoomUser(client),
	}
	log.Println("notify message: ", message)
	room.broadcastToClientsInRoom(message.encode())
}

func (room *Room) notifyClientLeft(client *Client) {
	message := &Message{
		Action:  LeaveRoomAction,
		Target:  room.name,
		Message: fmt.Sprintf(leaveMessage, client.GetName()),
		User:    room.ListRoomUser(client),
	}
	fmt.Println(message.User)
	log.Println("notify message: ", message)
	room.broadcastToClientsInRoom(message.encode())
}

func (room *Room) ListRoomUser(client *Client) []string {
	NonRepeatedUser := make(map[string]bool)
	for user := range room.clients {
		// fmt.Printf("user %s in room %s\n", user.Name, room.name)
		if _, ok := NonRepeatedUser[user.Name]; !ok {
			NonRepeatedUser[user.Name] = true
		}
	}

	result := []string{}
	for nonrepeateduser := range NonRepeatedUser {
		result = append(result, nonrepeateduser)
	}
	// fmt.Println("result: ", result)
	return result
}

func (room *Room) GetName() string {
	return room.name
}
