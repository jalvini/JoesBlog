package handlers

import (
	"Site1/models"
	"html/template"
	"net/http"
)

func RegisterHandler(write http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/user/register.html", "views/shared/header.html", "views/shared/footer.html")

	if err != nil {
		panic(err.Error())
	}

	template.Execute(write, "")
}

func RegisterPostHandler(write http.ResponseWriter, request *http.Request) {
	var message string

	request.ParseForm()

	username := request.Form.Get("username")
	password := request.Form.Get("password")
	firstName := request.Form.Get("first_name")
	lastName := request.Form.Get("last_name")

	registerTemplate, err := template.ParseFiles("views/user/register.html", "views/shared/header.html", "views/shared/footer.html")

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
