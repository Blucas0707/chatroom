package server

import (
	"log"
	chat "main/module/chat/controller"
	"main/module/user/controller"
	"os"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// TODO: wait websocket: https://dev.to/jeroendk/building-a-simple-chat-application-with-websockets-in-go-and-vue-js-gao
//https://dev.to/jeroendk/multi-room-chat-application-with-websockets-in-go-and-vue-js-part-2-3la8
func InitServer() *echo.Echo {
	//InitServer
	server := echo.New()
	//setting static: 例如：static目录下存在js/index.js文件， 则这个js的url为：/static/js/index.js
	server.Static("/static", "static")
	server.GET("/", func(c echo.Context) error {
		return c.File("templates/index.html")
	})
	//Session call
	sessionStore := sessions.NewCookieStore([]byte(sessionKey))
	server.Use(session.Middleware(sessionStore))

	//Register
	server.GET("/register", func(c echo.Context) error {
		return c.File("templates/register.html")
	})

	//api Login
	server.PATCH("/api/user", userlogin)
	//api Register
	server.POST("/api/user", userregister)
	//api Login info
	server.GET("/api/user", userlogininfo)
	//api Logout
	server.DELETE("/api/user", userlogout)

	//api createRoom
	server.POST("/api/room", createRoom)
	//api getRoomList
	server.GET("/api/rooms", getRoomList)

	//chatroom

	roomserver := NewRoomServer()
	go roomserver.Run()
	server.GET("/chatroom", func(c echo.Context) error {
		return c.File("templates/chatroom.html")
	})
	server.GET("/chatroom/ws", func(c echo.Context) error {
		roomName := c.QueryParam("room")
		if len(roomName) != 0 {
			// go roomserver.CreateRoom(roomName)
			serveWs(roomserver, c)
		} else {
			log.Println("No room name")
		}
		return nil
	})
	// fmt.Println(Server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	path := fmt.Sprintf(":%s", port)
	server.Logger.Fatal(server.Start(path))
	return server
}

var sessionKey string = os.Getenv("SESSION_KEY")

func userlogout(c echo.Context) error {
	type responseMessage struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	//delete session
	err := deleteSession(c)
	fmt.Println("logout:session save done")
	if err != nil {
		log.Println("failed to delete session:", err)
		response := &responseMessage{
			Error:   true,
			Message: "Internal Server Error",
		}
		return c.JSONPretty(http.StatusInternalServerError, response, "    ")
	}

	response := &responseMessage{
		Error:   false,
		Message: "Log out successfully",
	}
	getSession(c)
	return c.JSONPretty(http.StatusInternalServerError, response, "    ")
}

func userlogininfo(c echo.Context) error {
	type userdata struct {
		Username  string `json:"username,omitempty"`
		Useremail string `json:"useremail,omitempty"`
	}
	type responseMessage struct {
		Data userdata `json:"data"`
	}

	type errormessage struct {
		Data []int `json:"data"`
	}

	//get session
	name, email, password := getSession(c)
	log.Printf("session name:%s,email:%s,password:%s", name, email, password)
	if name != "-1" && email != "-1" && password != "-1" {

		userdata := userdata{
			Username:  name,
			Useremail: email,
		}

		response := responseMessage{
			Data: userdata,
		}
		return c.JSONPretty(http.StatusOK, response, "    ")
	} else {
		response := errormessage{
			Data: nil,
		}
		return c.JSONPretty(http.StatusOK, response, "    ")
	}
}

func userlogin(c echo.Context) error {

	//get json request
	json_map := make(map[string]interface{})
	if err := c.Bind(&json_map); err != nil {
		return err
	}

	useremail := fmt.Sprintf("%v", json_map["email"])
	userpassword := fmt.Sprintf("%v", json_map["password"])

	result, username := controller.Login(useremail, userpassword)

	//save session
	if result.ErrorExist == false {
		err := saveSession(c, username, useremail, userpassword)
		if err != nil {
			log.Printf("session error:", err)
		}
	}

	log.Println(result)
	return c.JSONPretty(http.StatusOK, result, "    ")
}

func saveSession(c echo.Context, name, email, password string) error {
	sess, err := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	//Save session
	sess.Values["username"] = name
	sess.Values["useremail"] = email
	sess.Values["userpassword"] = password
	sess.Save(c.Request(), c.Response())
	return err
}

// TODO: wait to test logout function
func deleteSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values["username"] = ""
	sess.Save(c.Request(), c.Response())
	return err
}

func getSession(c echo.Context) (string, string, string) {
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("getSession err")
	}
	if sess.Values["username"] == nil {
		return "-1", "-1", "-1"
	} else {
		//get session
		name := fmt.Sprintf("%v", sess.Values["username"])
		email := fmt.Sprintf("%v", sess.Values["useremail"])
		password := fmt.Sprintf("%v", sess.Values["userpassword"])
		return name, email, password
	}
}

//TODO: wait api https://app.swaggerhub.com/apis-docs/padax/taipei-trip/1.0.0?loggedInWithGitHub=true#/%E4%BD%BF%E7%94%A8%E8%80%85/patch_api_user
func userregister(c echo.Context) error {
	//get json request
	json_map := make(map[string]interface{})
	if err := c.Bind(&json_map); err != nil {
		return err
	}
	username := fmt.Sprintf("%v", json_map["name"])
	useremail := fmt.Sprintf("%v", json_map["email"])
	userpassword := fmt.Sprintf("%v", json_map["password"])
	passwordConfirm := fmt.Sprintf("%v", json_map["repassword"])
	result := controller.Register(username, useremail, userpassword, passwordConfirm)
	return c.JSONPretty(http.StatusOK, result, "    ")
}

func createRoom(c echo.Context) error {
	owner, _, _ := getSession(c)
	if owner == "-1" {
		return c.String(http.StatusOK, "please login first")
	}
	//get json request
	json_map := make(map[string]interface{})
	if err := c.Bind(&json_map); err != nil {
		return err
	}
	chatroomName := fmt.Sprintf("%v", json_map["roomname"])
	result := chat.CreateRoom(chatroomName, owner)
	return c.JSONPretty(http.StatusOK, result, "    ")
}

func getRoomList(c echo.Context) error {
	owner, _, _ := getSession(c)
	if owner == "-1" {
		return c.String(http.StatusOK, "please login first")
	}
	page := c.QueryParam("page")
	pageInt, _ := strconv.Atoi(page)
	result := chat.GetRoom(pageInt)
	return c.JSONPretty(http.StatusOK, result, "    ")
}
