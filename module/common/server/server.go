package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitServer() {
	//InitServer
	Server := echo.New()
	Server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	fmt.Println(Server)
	Server.Logger.Fatal(Server.Start(":1323"))
}
