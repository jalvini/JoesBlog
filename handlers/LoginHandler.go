package handlers

import (
	"Site1/helpers"
	"Site1/models"
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(write http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/user/login.html", "views/shared/header.html", "views/shared/footer.html")

	if err != nil {
		fmt.Println(write, "Failed To Load Template")
		panic(err.Error())
	}

	template.Execute(write, "This is a template")
}

func LoginPostHandler(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	loginCheck := models.UserGuard(username, password)

	redirectTarget := "/login"

	if loginCheck == true {
		helpers.SetCookie(username, write)
		//redirectTarget = "/index"
	} else {
		redirectTarget = "/register"
	}

	http.Redirect(write, request, redirectTarget, 302)

	fmt.Fprintln(write, bool(bool(loginCheck)))
}
