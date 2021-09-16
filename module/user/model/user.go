package model

import (
	"chatroom/module/common/db_module"
	. "chatroom/module/common/db_module"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"unicode"
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

	// //Check password format:至少 8 碼，需有英文大小寫與數字混用，至少要有一個英文字母與數字。
	wg.Add(1)
	go func(c chan bool) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult := Validation(password)
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

func Validation(password string) bool {
	var (
		isUpper  = false
		isLower  = false
		isNumber = false
	)
	if len(password) < 8 {
		return false
	}
	for _, s := range password {
		switch {
		case unicode.IsUpper(s):
			isUpper = true
		case unicode.IsLower(s):
			isLower = true
		case unicode.IsNumber(s):
			isNumber = true
		}
	}
	if (isUpper && isLower) && isNumber {
		return true
	}
	return false
}
