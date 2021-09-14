package db_module

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//InitDB
func InitDB() (*sql.DB, error) {
	//Create connection pool
	fmt.Println("Preparing to create a connection pool")
	db, err := sql.Open("mysql", "root:root1234@tcp(localhost)/chatroom?charset=utf8mb4,utf8&parseTime=true")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err != nil { // Init Error
		return nil, err
	} else {
		//Ping
		fmt.Println("Start to ping")
		if pingErr := db.Ping(); pingErr != nil {
			return nil, pingErr
		} else {
			fmt.Println("Ping successfully")
			return db, nil
		}
	}
}

//Insert
