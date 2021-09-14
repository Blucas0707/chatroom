package model

import (
	// "chatroom/module/common/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userAccount struct {
	name     string `json:name`
	email    string `json:email`
	password string `json:password`
}

var user = &userAccount{}

// func init() {
// 	server.InitServer()
// }

//Get user register data
func GetRegister(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
