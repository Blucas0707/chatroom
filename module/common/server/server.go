package server

import (
	"chatroom/module/user/controller"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitServer() *echo.Echo {
	//InitServer
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	//Register
	server.GET("/register", func(c echo.Context) error {
		return c.File("templates/register.html")
	})

	//api Login
	server.PUT("/api/user", userlogin)
	//api Register
	server.POST("/api/user", userregister)
	// fmt.Println(Server)
	server.Logger.Fatal(server.Start(":1323"))
	return server
}

type responseMessage struct {
	Status  string `json:error`
	Message string `json:message`
}

func userlogin(c echo.Context) error {
	// Get email and password
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+" email:"+email)
}

func userregister(c echo.Context) error {
	//Get name, email and password
	name := c.FormValue("name")
	fmt.Println("name:", name)
	email := c.FormValue("email")
	password := c.FormValue("password")
	passwordConfirm := c.FormValue("password-confirm")
	res := new(responseMessage)
	if err := c.Bind(res); err != nil {
		return err
	}
	// password reconfirm
	if password != passwordConfirm {
		res.Status = "false"
		res.Message = "not the same"
		fmt.Println(res)
		return c.JSON(http.StatusCreated, res)
	}
	fmt.Println("prepare check")
	// check user name taken
	// result := model.CheckUserRegister(name, email, password)
	result := controller.Register(name, email, password)
	fmt.Println(result)
	return c.String(http.StatusOK, "name:"+name+" email:"+email+" password:"+password+" passwordConfirm:"+passwordConfirm)
}
