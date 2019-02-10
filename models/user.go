package models

import (
	"Site1/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	pwd := getPwd(password)

	defer db.Close()

	err := db.QueryRow("SELECT id, username, password, first_name, last_name FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Joe")
	pwdMatch := comparePasswords(user.Password, pwd)

	return pwdMatch
}

func CreateUser(username string, password string, firstName string, lastName string) string {
	db := database.DbInit()

	pwd := getPwd(password)
	hash := hashAndSalt(pwd)

	insert, err := db.Query("INSERT INTO users VALUES (?, ?, ?, ?, ? )", nil, username, hash, firstName, lastName)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	return "User Created Successfully"
}

func getPwd(pwd string) []byte {
	_, err := fmt.Scan(pwd)
	if err != nil {
		log.Println(err)
	}

	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
