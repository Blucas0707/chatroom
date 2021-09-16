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
	//setting static: 例如：static目录下存在js/index.js文件， 则这个js的url为：/static/js/index.js
	server.Static("/static", "static")
	server.GET("/", func(c echo.Context) error {
		return c.File("templates/index.html")
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
	//get json request
	json_map := make(map[string]interface{})
	if err := c.Bind(&json_map); err != nil {
		return err
	}
	json_name := fmt.Sprintf("%v", json_map["name"])
	fmt.Println(json_name)
	return c.String(http.StatusOK, json_name)

	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	// if err != nil {
	// 	return err
	// } else {
	// 	//json_map has the JSON Payload decoded into a map
	// 	json_name := json_map["name"]
	// 	json_email := json_map["email"]
	// 	json_password := json_map["password"]
	// 	fmt.Printf(json_name, json_email, json_password)
	// }

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

	result := controller.Register(name, email, password, passwordConfirm)
	fmt.Println(result)
	return c.String(http.StatusOK, "name:"+name+" email:"+email+" password:"+password+" passwordConfirm:"+passwordConfirm)
}
