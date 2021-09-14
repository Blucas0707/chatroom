package db_module

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func InitEnv() (err error) {
	const projectDirName = "chatroom"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	loadingErr := godotenv.Load(string(rootPath) + `/module/common/db_module/.env`)
	if loadingErr != nil {
		log.Fatal("Error loading .env file", loadingErr)
		return loadingErr
	}
	return nil
}

//InitDB
func InitDB() (*sql.DB, error) {
	//Load .env
	InitEnv()

	//Create connection pool
	fmt.Println("Preparing to create a connection pool")
	sql_connection_info := os.Getenv("SQL_CONNECTION_INFO")
	db, err := sql.Open("mysql", sql_connection_info)
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
