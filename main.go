package main

import (
	. "chatroom/module/common/server"
)

// . "chatroom/module/common/server"

type user struct {
	username     string
	useremail    string
	userpassword string
}

func main() {
	InitServer()
	//InitDB
	// db, err := db_module.InitDB()
	// if err != nil {
	// 	fmt.Println("DB init error:", err)
	// 	return
	// }
	// server.InitServer()

	// testEmail := "test@test.comsss"
	// checkEmailResult, err := db_module.CheckEmailisNotExisted(db, testEmail)
	// fmt.Println(checkEmailResult, err)

	// testName := "test123"
	// checkNameResult, err := db_module.CheckNameisNotExisted(db, testName)
	// fmt.Println(checkNameResult, err)

	// testuser := user{
	// 	username:     "testuser",
	// 	useremail:    "test2@test.com",
	// 	userpassword: "test",
	// }
	// registerResult, err := db_module.UserRegister(db, testuser.username, testuser.useremail, testuser.userpassword)
	// fmt.Println(registerResult, err)

	// testuseremail := "test@test.com"
	// testpassword := "test"
	// LoginResult, err := db_module.CheckUserLogin(db, testuseremail, testpassword)
	// fmt.Println(LoginResult, err)
}
