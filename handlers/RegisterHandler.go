package handlers

import (
	"Site1/models"
	"fmt"
	"html/template"
	"net/http"
)

func RegisterHandler(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "text/html")

	template, err := template.ParseFiles("views/user/register.gohtml", "views/shared/header.gohtml", "views/shared/footer.gohtml")

	if err != nil {
		fmt.Println(write, "Failed To Load Template")
		panic(err.Error())
	}

	tempErr := template.Execute(write, "John Smith")
	if tempErr != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterPostHandler(write http.ResponseWriter, request *http.Request) {
	var message string

	request.ParseForm()

	username := request.Form.Get("username")
	password := request.Form.Get("password")
	firstName := request.Form.Get("first_name")
	lastName := request.Form.Get("last_name")

	registerTemplate, err := template.ParseFiles("views/user/register.gohtml", "views/shared/header.gohtml", "views/shared/footer.gohtml")

	if err != nil {
		panic(err.Error())
	}

	if len(username) <= 3 || len(password) <= 5 || len(firstName) == 0 || len(lastName) == 0 {
		message = "You Did Something Wrong. Username Must Be at Least 4 Chars, Password Must Be 5 Chars and Name Cant Be Empty"
	} else {
		models.CreateUser(username, password, firstName, lastName)
		message = "User Successfully Created"
	}

	registerTemplate.Execute(write, message)
}
