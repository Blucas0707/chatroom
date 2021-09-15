package db_module

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

var DB *sql.DB

//InitDB
func InitDB() (*sql.DB, error) {
	//Create connection pool
	fmt.Println("Preparing to create a connection pool")
	sql_connection_info := os.Getenv("SQL_CONNECTION_INFO")
	DB, err := sql.Open("mysql", sql_connection_info)
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	if err != nil {
		fmt.Println("error occur when connect to sql:", err)
		return nil, err
	} else {
		fmt.Println("connect to DB successfully")
		return DB, nil
	}
}

type user struct {
	id    string
	name  string
	email string
}

func CheckEmail(email string) bool {
	sql, err := DB.Prepare("SELECT count(*) from userinfo")
	if err != nil {
		log.Println("checkEmail query stmt error: ", err)
		return false
	}
	rows, err := sql.Query(email)
	if err != nil {
		log.Println("checkEmail query error: ", err)
		return false
	}
	fmt.Println(rows)
	return true
}
