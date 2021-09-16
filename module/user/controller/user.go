package controller

import (
	. "chatroom/module/common/db_module"
	// "chatroom/module/common/user/model"
	"chatroom/module/user/model"
	"fmt"
	"log"
)

func Login(useremail, password string) []string {
	result, err := model.UserLogin(useremail, password)
	if err != nil {
		log.Printf("User Login Server Error: ", err)
		return []string{"-1"}
	}
	return result
}

//Before register, check if name and email are taken
func Register(name, email, password, passwordConfirm string) string {
	db, _ := InitDB()
	allowedRegister := model.CheckUserRegister(name, email, password, passwordConfirm)
	if allowedRegister {
		fmt.Println("Register Allowed")
		fmt.Printf("Registering...")
		registerStatus, _ := UserRegister(db, name, email, password)
		if registerStatus {
			fmt.Printf("Register success!")
			return "Register success"
		} else {
			fmt.Printf("Register fail!")
			return "Register fail"
		}
	} else {
		return "Register Not Allowed"
	}

}
