package db_module

import (
	"database/sql"
	"fmt"
	"log"
	. "main/module/common/error"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

// var db *sql.DB

const (
	ConnMaxLifetime int64 = 3
	MaxOpenConns    int   = 10
	MaxIdleConns    int   = 10
)

//InitDB
func InitDB() (*sql.DB, error) {
	//Create connection pool
	// fmt.Println("Preparing to create a connection pool")
	sql_connection_info := os.Getenv("SQL_CONNECTION_INFO")
	db, err := sql.Open("mysql", sql_connection_info)
	if err != nil {
		fmt.Println("error occur when connect to sql:", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("error occur when ping sql_db:", err)
		return nil, err
	}

	// fmt.Println("Set DB ConnMaxLifetime: ", ConnMaxLifetime)
	db.SetConnMaxLifetime(time.Minute * time.Duration(ConnMaxLifetime))
	// fmt.Println("Set DB MaxOpenConns: ", MaxOpenConns)
	db.SetMaxOpenConns(MaxOpenConns)
	// fmt.Println("Set DB MaxIdleConns: ", MaxIdleConns)
	db.SetMaxIdleConns(MaxIdleConns)
	// fmt.Println("connect to DB successfully")
	return db, nil

}

// Check email is taken
func CheckEmailisNotExisted(db *sql.DB, email string) (bool, error) {
	sel, err := db.Prepare("select count(*) from userinfo where user_email = ?")
	if err != nil {
		log.Println("sel error:", err)
		return false, err
	}
	defer sel.Close()

	rows, err := sel.Query(email)
	if CheckError(err) {
		return false, err
	}

	defer rows.Close()
	var count int
	if rows.Next() {
		err := rows.Scan(&count)
		if CheckError(err) {
			return false, err
		}
		fmt.Println("count:", count)
	} else {
		log.Printf("email %s not found\n", email)
	}

	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

//check name is taken
func CheckNameisNotExisted(db *sql.DB, name string) (bool, error) {
	sel, err := db.Prepare("select count(*) from userinfo where user_name = ?")
	if err != nil {
		log.Println("sel error:", err)
		return false, err
	}
	defer sel.Close()

	rows, err := sel.Query(name)
	if CheckError(err) {
		return false, err
	}

	defer rows.Close()
	var count int
	if rows.Next() {
		err := rows.Scan(&count)
		if CheckError(err) {
			return false, err
		}
		fmt.Println("count:", count)
	} else {
		log.Printf("name %s not found\n", name)
	}
	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

//TODO: Encrypt password
//user register
func UserRegister(db *sql.DB, username, useremail, userpassword string) (bool, error) {
	ins, err := db.Prepare("insert into userinfo(user_name, user_email, user_password) values(?, ?, ?)")
	if err != nil {
		log.Println("sel error:", err)
		return false, err
	}
	defer ins.Close()

	result, err := ins.Exec(username, useremail, userpassword)
	if CheckError(err) {
		return false, err
	}
	defer ins.Close()
	id, err := result.LastInsertId()
	if CheckError(err) {
		return false, err
	}
	fmt.Printf("LastInsertId: %d\n", id)
	return true, nil
}

//user login
func UserLogin(db *sql.DB, useremail string, userpassword string) (int, string, error) {
	sel, err := db.Prepare("SELECT count(*), user_name from userinfo where user_email = ? and user_password = ? group by user_name limit 1")
	if err != nil {
		log.Printf("CheckUserLogin error: %v\n", err)
		return -1, "", err
	}
	defer sel.Close()

	rows, err := sel.Query(useremail, userpassword)
	if err != nil {
		log.Printf("CheckUserLogin query error:%v \n", err)
		return -1, "", err
	}
	defer rows.Close()

	user_count := 0
	user_name := ""
	if rows.Next() {
		err := rows.Scan(&user_count, &user_name)
		if err != nil {
			log.Printf("CheckUserLogin rows error: %v \n", err)
			return user_count, "", err
		}
	} else {
		log.Printf("user %s not found\n", useremail)
		return user_count, "", err
	}

	if user_count == 1 {
		log.Printf("user %s found\n", useremail)
	} else {
		return user_count, user_name, nil
	}
	return user_count, user_name, nil
}

// chatroom
// check room name is not existed
func CheckRoomNameisNotExisted(db *sql.DB, name string) (bool, error) {
	sel, err := db.Prepare("select count(*) from chatroom_list  where name  = ?")
	if err != nil {
		log.Println("sel error:", err)
		return false, err
	}
	defer sel.Close()

	rows, err := sel.Query(name)
	if CheckError(err) {
		return false, err
	}

	defer rows.Close()
	var count int
	if rows.Next() {
		err := rows.Scan(&count)
		if CheckError(err) {
			return false, err
		}
		fmt.Println("count:", count)
	} else {
		log.Printf("room name %s not found\n", name)
	}
	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

//user register
func CreateChatRoom(db *sql.DB, chatroomName, owner string) (bool, error) {
	ins, err := db.Prepare("insert into chatroom_list (name , owner) values(?, ?)")
	if err != nil {
		log.Println("sel error:", err)
		return false, err
	}
	defer ins.Close()

	result, err := ins.Exec(chatroomName, owner)
	if CheckError(err) {
		return false, err
	}
	defer ins.Close()
	id, err := result.LastInsertId()
	if CheckError(err) {
		return false, err
	}
	fmt.Printf("LastInsertId: %d\n", id)
	return true, nil
}

// get Room list
func GetRoomList(db *sql.DB, page int) ([]string, []string) {
	sel, err := db.Prepare("select name,owner from chatroom_list order by id DESC limit ?,10")
	if err != nil {
		log.Println("sel error:", err)
	}
	defer sel.Close()
	page = page * 10
	rows, err := sel.Query(page)
	if err != nil {
		log.Println("sel error:", err)
	}

	defer rows.Close()
	var chatrooms []string
	var owners []string

	for rows.Next() {
		var room, owner string
		if err := rows.Scan(&room, &owner); err != nil {
			log.Fatal(err)
		} else {
			chatrooms = append(chatrooms, room)
			owners = append(owners, owner)
		}

	}
	return chatrooms, owners
}
