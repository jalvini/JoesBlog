package models

import (
	"Site1/database"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func UserGuard(username string, password string) bool {
	var user User
	db := database.DbInit()

	defer db.Close()

	err := db.QueryRow("SELECT id, username, password, first_name, last_name FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)

	if err != nil {
		return false
	}

	return true
}

func CreateUser(username string, password string, firstName string, lastName string) string {
	db := database.DbInit()

	insert, err := db.Query("INSERT INTO users VALUES (?, ?, ?, ?, ? )", nil, username, password, firstName, lastName)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	return "User Created Successfully"
}
