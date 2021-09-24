package server

// TODO: https://github.com/jeroendk/go-vuejs-chat/blob/v2.0/chatServer.go

// a server which controll rooms
type RoomServer struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Registered rooms
	rooms map[*Room]bool
}

func NewRoomServer() *RoomServer {
	return &RoomServer{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[*Room]bool),
	}
}

func (roomserver *RoomServer) Run() {
	for {
		select {
		case client := <-roomserver.register:
			roomserver.registerClient(client)
		case client := <-roomserver.unregister:
			roomserver.unregisterClient(client)
		case message := <-roomserver.broadcast:
			roomserver.broadcastToClients(message)
		}

	}
}

func (roomserver *RoomServer) registerClient(client *Client) {
	// roomserver.notifyClientJoined(client)
	roomserver.clients[client] = true
}

func (roomserver *RoomServer) unregisterClient(client *Client) {
	if _, ok := roomserver.clients[client]; ok {
		delete(roomserver.clients, client)
		// roomserver.notifyClientLeft(client)
	}
}

func (roomserver *RoomServer) broadcastToClients(message []byte) {
	for client := range roomserver.clients {
		client.send <- message
	}
}

func (roomserver *RoomServer) findRoombyName(name string) *Room {
	var foundRoom *Room
	for room := range roomserver.rooms {
		if room.GetName() == name {
			foundRoom = room
			break
		}
	}
	return foundRoom
}

func (roomserver *RoomServer) CreateRoom(name string) *Room {
	room := NewRoom(name)
	go room.RunRoom()
	roomserver.rooms[room] = true
	return room
}
