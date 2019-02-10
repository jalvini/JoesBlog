package handlers

import (
	"Site1/models"
	"fmt"
	"html/template"
	"net/http"
)

func RegisterHandler(write http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/user/register.html", "views/shared/header.html")

	if err != nil {
		fmt.Println(write, "Failed To Load Template")
		panic(err.Error())
	}

	template.Execute(write, "This is a template")
}

func RegisterPostHandler(write http.ResponseWriter, request *http.Request) {
	registerTemplate, err := template.ParseFiles("views/user/register.html", "views/shared/header.html")

	if err != nil {
		panic(err.Error())
	}

	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	firstName := request.Form.Get("first_name")
	lastName := request.Form.Get("last_name")
	var message string

	if len(username) == 0 || len(password) == 0 || len(firstName) == 0 || len(lastName) == 0 {
		message = "You Did Something Wrong"
	} else {
		models.CreateUser(username, password, firstName, lastName)
		message = "User Successfully Created"
	}

	registerTemplate.Execute(write, message)
}
