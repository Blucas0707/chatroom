package model

import (
	"chatroom/module/common/db_module"
	. "chatroom/module/common/db_module"
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// type userAccount struct {
// 	id       string `json:id`
// 	name     string `json:name`
// 	email    string `json:email`
// 	password string `json:password`
// }

// var userData = &userAccount{}

//TODO: implement responseMessage struct
// var responseMessage = map[string]struct{}{
// 	"1": {
// 		err:       "true",
// 		"message": "Name is taken",
// 	},
// }

var db *sql.DB

func init() {
	db, _ = db_module.InitDB()
}

func UserLogin(email, password string) ([]string, error) {
	result, err := db_module.CheckUserLogin(db, email, password)
	if err != nil {
		log.Println("CheckUserLogin error:", err)
		return []string{"0"}, err
	}
	return result, nil
}

func CheckUserRegister(name, email, password string) bool {
	registerAllowed := make(chan bool, 1)
	defer close(registerAllowed)
	registerAllowed <- true
	wg := sync.WaitGroup{}

	// Check Name is Not Existed
	wg.Add(1)
	go func(c chan bool) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult, _ := CheckNameisNotExisted(db, name)
				fmt.Println(checkResult)
				c <- checkResult
				wg.Done()
			} else {
				c <- false
				wg.Done()
			}
		}
	}(registerAllowed)

	// Check Email is Not Existed
	wg.Add(1)
	go func(c chan bool) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult, _ := CheckEmailisNotExisted(db, name)
				fmt.Println(checkResult)
				c <- checkResult
				wg.Done()
			} else {
				c <- false
				wg.Done()
			}
		}
	}(registerAllowed)

	wg.Wait()

	result := <-registerAllowed
	return result
}
