package main

import (
	"chatroom/module/common/db_module"
	"chatroom/module/common/server"
	"chatroom/module/user/model"

	"github.com/labstack/echo/v4"
)

func main() {
	//InitDB
	db_module.InitDB()
	server.InitServer()
	c := echo.Context
	model.GetRegister(c)
}

// func getUser(c echo.Context) error {
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }
