package model

import (
	"database/sql"
	"fmt"
	"log"
	"main/module/common/db_module"
	"regexp"
	"sync"
	"unicode"
)

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
		ErrorMessage: "Register successed, please sign in again",
	},
	1: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Username is existed, please try again",
	},
	2: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Useremail is existed, please try again",
	},
	3: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Password confirmation failed, please try again",
	},
	4: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Email format error, please try again",
	},
	5: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Password format error, please try again",
	},
	6: {
		ErrorCode:    200,
		ErrorExist:   false,
		ErrorMessage: "Login successed",
	},
	7: {
		ErrorCode:    400,
		ErrorExist:   true,
		ErrorMessage: "Email or password not found or error, please check",
	},
	8: {
		ErrorCode:    500,
		ErrorExist:   true,
		ErrorMessage: "Internal server error",
	},
}

var db *sql.DB

func init() {
	db, _ = db_module.InitDB()
}

//user register
func CreateUser(username, useremail, userpassword string) (bool, error) {
	ins, err := db.Prepare("insert into userinfo(user_name, user_email, user_password) values(?, ?, ?)")
	if err != nil {
		log.Println("sel error:", err)
		return false, err
	}
	defer ins.Close()

	result, err := ins.Exec(username, useremail, userpassword)
	if err != nil {
		return false, err
	}
	defer ins.Close()
	id, err := result.LastInsertId()
	if err != nil {
		return false, err
	}
	fmt.Printf("LastInsertId: %d\n", id)
	return true, nil
}

func CheckUserLogin(email, password string) (Message, string) {
	result, username, _ := db_module.UserLogin(db, email, password)
	errorcode := 0

	if result == 0 {
		errorcode = 7
	} else if result == -1 {
		errorcode = 8
	} else if result == 1 {
		errorcode = 6
	}

	errorMessage := responseMessage[errorcode]
	return errorMessage, username
}

func CheckUserRegister(name, email, password, passwordConfirm string) Message {
	registerAllowed := make(chan bool, 1)
	defer close(registerAllowed)
	registerAllowed <- true

	ErrorMessageCodechan := make(chan int, 1)
	ErrorMessageCodechan <- 0
	wg := sync.WaitGroup{}

	//Check Email is valid
	wg.Add(1)
	go func(c chan bool, e chan int) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult := isEmailValid(email)
				fmt.Println("email format check:", checkResult)
				c <- checkResult

				if !checkResult {
					<-e
					e <- 4
				}

				wg.Done()
			} else {
				c <- false
				wg.Done()
			}
		}
	}(registerAllowed, ErrorMessageCodechan)

	// //Check password format:至少 8 碼，需有英文大小寫與數字混用，至少要有一個英文字母與數字。
	wg.Add(1)
	go func(c chan bool, e chan int) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult := isPasswordValid(password, passwordConfirm)
				fmt.Println("password check:", checkResult)
				c <- checkResult

				if !checkResult {
					<-e
					e <- 5
				}

				wg.Done()
			} else {
				c <- false
				wg.Done()
			}
		}
	}(registerAllowed, ErrorMessageCodechan)

	// Check Name is Existed
	wg.Add(1)
	go func(c chan bool, e chan int) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult, _ := db_module.CheckNameisNotExisted(db, name)
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
	}(registerAllowed, ErrorMessageCodechan)

	// Check Email is Existed
	wg.Add(1)
	go func(c chan bool, e chan int) {
		select {
		case iscontinue := <-c:
			if iscontinue {
				checkResult, _ := db_module.CheckEmailisNotExisted(db, name)
				fmt.Println("email check:", checkResult)
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
	}(registerAllowed, ErrorMessageCodechan)

	wg.Wait()

	ErrorMessageCode := <-ErrorMessageCodechan
	result := responseMessage[ErrorMessageCode]
	fmt.Println(ErrorMessageCode)
	return result
}

func isPasswordValid(password, passwordConfirm string) bool {
	var (
		isUpper  = false
		isLower  = false
		isNumber = false
	)
	// check password & re-enter password same
	if password != passwordConfirm {
		return false
	}
	// check password length >= 8
	if len(password) < 8 {
		return false
	}
	// check including number, Upper and Lower alphabelt
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

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
