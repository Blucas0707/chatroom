package main

import (
	"chatroom/module/common/db_module"
	"database/sql"
	"fmt"
)

func main() {
	//InitDB
	DB, _ := db_module.InitDB()
	// server.InitServer()

	// testEmail := "123@123.com"
	// db_module.CheckEmail(testEmail)
	CheckEmail(DB, 1)

	// c := echo.Context
	// model.GetRegister(c)

}

// func getUser(c echo.Context) error {
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

type Member struct {
	id int
}

//TODO
func CheckEmail(db *sql.DB, id int) (*Member, error) {
	mem := &Member{}
	err := db.QueryRow("select count(*) from userinfo where user_id = ?", id).Scan(&mem.id)
	fmt.Println(mem)
	return mem, err
}
