package controller

import (
	"main/module/chat/model"
	. "main/module/common/db_module"
)

func CreateRoom(chatroomName, owner string) model.Message {
	db, _ := InitDB()
	ErrorMessage := model.CheckRoomName(chatroomName)
	if !ErrorMessage.ErrorExist {
		createStatus, _ := CreateChatRoom(db, chatroomName, owner)
		if !createStatus {
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

func GetRoom(page int) model.ChatroomList {
	result := model.GetRoomData(page)
	return result
}
