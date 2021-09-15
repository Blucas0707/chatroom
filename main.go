package main

import (
	"chatroom/module/common/db_module"
	"database/sql"
)

func main() {
	//InitDB
	DB, err := db_module.InitDB()
	// server.InitServer()

	// testEmail := "123@123.com"
	// db_module.CheckEmail(testEmail)
	CheckEmail

	// c := echo.Context
	// model.GetRegister(c)

}

// func getUser(c echo.Context) error {
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

//TODO
func CheckEmail(db *sql.DB, id int) (int bool) {
	var uid int
	err := db.QueryRow("select count(*) from userinfo where user_id = ?", id).Scan(&uid)
	return uid, err
}
