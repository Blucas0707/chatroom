package model

import (
	"chatroom/module/common/server"
)

type userAccount struct {
	name     string
	email    string
	password string
}

var user = &userAccount{}

func init() {
	server.InitServer()
}

//Get user register data
func GetRegister() {
	Server.GET("/users/:id")
}
