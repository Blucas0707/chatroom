package main

import (
	"chatroom/module/common/server"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	server.InitServer()
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
