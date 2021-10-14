package model

import (
	"database/sql"
	"fmt"
	"main/module/common/db_module"
	. "main/module/common/db_module"
	"sync"
)

type ChatroomList struct {
	Page int        `json:page`
	Data []chatroom `json:"data"`
}

type chatroom struct {
	Name  string `json:"chatroomName"`
	Owner string `json:"owner"`
}

type Message struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorExist   bool   `json:"errorExist"`
	ErrorMessage string `json:"errorMessage"`
}

//Define responseMessage struct
var responseMessage = map[int]Message{
	0: {
		ErrorCode:    200,
		ErrorExist:   false,
		ErrorMessage: "Create successed",
	},
	1: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Name existed",
	},
	2: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Name length < 4",
	},
}

var db *sql.DB

func init() {
	db, _ = db_module.InitDB()
}

func CheckRoomName(chatroomName string) Message {

	createAllowed := make(chan bool, 1)
	defer close(createAllowed)
	createAllowed <- true

	ErrorMessageCodechan := make(chan int, 1)
	ErrorMessageCodechan <- 0
	wg := sync.WaitGroup{}

	//Check Name Length is valid (>=4)
	wg.Add(1)
	go func(c chan bool, e chan int) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult := func(chatroomName string) bool {
					if len(chatroomName) < 4 {
						return false
					} else {
						return true
					}
				}(chatroomName)
				fmt.Println("name format check:", checkResult)
				c <- checkResult

				if !checkResult {
					<-e
					e <- 2
				}

				wg.Done()
			} else {
				c <- false
				wg.Done()
			}
		}
	}(createAllowed, ErrorMessageCodechan)

	// Check Name is Existed
	wg.Add(1)
	go func(c chan bool, e chan int) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult, _ := CheckRoomNameisNotExisted(db, chatroomName)
				fmt.Println("name check:", checkResult)
				c <- checkResult

				if !checkResult {
					<-e
					e <- 1
				}

				wg.Done()
			} else {
				c <- false
				wg.Done()
			}
		}
	}(createAllowed, ErrorMessageCodechan)

	wg.Wait()

	ErrorMessageCode := <-ErrorMessageCodechan
	result := responseMessage[ErrorMessageCode]
	fmt.Println(ErrorMessageCode)
	return result
}

func GetRoomData(page int) ChatroomList {
	rooms, owners := GetRoomList(db, page)
	roomlist := ChatroomList{
		Page: page,
		Data: []chatroom{},
	}
	for i, _ := range rooms {
		roomlist.Data = append(roomlist.Data, chatroom{
			Name:  rooms[i],
			Owner: owners[i],
		})
	}
	// fmt.Println(roomlist)
	return roomlist
}
