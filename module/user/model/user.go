package model

import (
	"chatroom/module/common/db_module"

	"github.com/labstack/echo/v4"
	"net/http"
)

type userAccount struct {
	id       string `json:id`
	name     string `json:name`
	email    string `json:email`
	password string `json:password`
}

var userData = &userAccount{}

// func init() {
// 	server.InitServer()
// }

//Get user data
func GetRegister(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
