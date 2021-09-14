package main

import (
	"fmt"
	"log"

	"chatroom/db_module"
)

func main() {
	//InitDB
	db, err := db_module.InitDB()
	if err != nil {
		log.Fatal("Initial db: ", err)
	}
	defer db.Close()

	type UserData struct {
		name     string
		email    string
		password string
	}

	inputData := UserData{
		name:     "user01",
		email:    "user01@example.com",
		password: "user01_password",
	}
	fmt.Println(inputData)
	// db_module.InsertDB(&inputData)

	//InitServer
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
