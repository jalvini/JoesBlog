package handlers

import (
	"Site1/helpers"
	"Site1/models"
	"html/template"
	"net/http"
)

func RegisterHandler(write http.ResponseWriter, request *http.Request) {
	checkUser := helpers.GetUserCookie(request)

	if len(checkUser) > 0 {
		redirectTarget := "/dashboard"
		http.Redirect(write, request, redirectTarget, 302)
	}

	write.Header().Set("Content-Type", "text/html")

	err := helpers.Template.ExecuteTemplate(write, "register.gohtml", "John Smith")
	if err != nil {
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
