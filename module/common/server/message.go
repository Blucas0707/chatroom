package server

import (
	"encoding/json"
	"log"
)

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"

type Message struct {
	Action  string   `json:"action"`
	Message string   `json:"message"`
	Target  string   `json:"target"`
	Sender  *Client  `json:"sender"`
	User    []string `json:"user"`
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println("encode err: ", err)
	}
	return json
}
