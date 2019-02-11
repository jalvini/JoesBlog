package models

import (
	"Site1/database"
	"Site1/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func UserGuard(username string, password string) bool {
	var user User
	db := database.DbInit()
	pwd := helpers.GetPwd(password)

	defer db.Close()

	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return false
	}

	pwdMatch := helpers.ComparePasswords(user.Password, pwd)

	return pwdMatch
}

func ShowUser(username string) User {
	var user User
	db := database.DbInit()

	defer db.Close()

	err := db.QueryRow("SELECT id, username, password, first_name, last_name FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)

	if err != nil {
		panic(err.Error())
	}

	return user
}

func CreateUser(username string, password string, firstName string, lastName string) string {
	db := database.DbInit()

	pwd := helpers.GetPwd(password)
	hash := helpers.HashAndSalt(pwd)

	insert, err := db.Query("INSERT INTO users VALUES (?, ?, ?, ?, ? )", nil, username, hash, firstName, lastName)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	return "User Created Successfully"
}
