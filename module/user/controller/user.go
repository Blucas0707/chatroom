package controller

import (
	. "main/module/common/db_module"
	// "chatroom/module/common/user/model"
	"main/module/user/model"
)

//Before register, check if name and email are taken
func Register(name, email, password, passwordConfirm string) model.Message {
	db, _ := InitDB()
	ErrorMessage := model.CheckUserRegister(name, email, password, passwordConfirm)
	if !ErrorMessage.ErrorExist {
		registerStatus, _ := UserRegister(db, name, email, password)
		if !registerStatus {
			internalError := model.Message{
				ErrorCode:    500,
				ErrorExist:   true,
				ErrorMessage: "Internal Error",
			}
			return internalError
		}
	}
	return ErrorMessage

}

func Login(email, password string) (model.Message, string) {
	ErrorMessage, username := model.CheckUserLogin(email, password)
	return ErrorMessage, username

}
