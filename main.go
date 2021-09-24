package main

import (
	. "main/module/common/server"
)

type user struct {
	username     string
	useremail    string
	userpassword string
}

func main() {
	InitServer()
}
