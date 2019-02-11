package handlers

import (
	"Site1/helpers"
	"Site1/models"
	"fmt"
	"net/http"
)

func LoginHandler(write http.ResponseWriter, request *http.Request) {
	checkUser := helpers.GetUserCookie(request)

	if len(checkUser) > 0 {
		redirectTarget := "/dashboard"
		http.Redirect(write, request, redirectTarget, 302)
	}

	write.Header().Set("Content-Type", "text/html")

	err := helpers.Template.ExecuteTemplate(write, "login.gohtml", "John Smith")
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPostHandler(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	loginCheck := models.UserGuard(username, password)

	redirectTarget := "/login"

	if loginCheck == true {
		helpers.SetUserCookie(username, write)
		//redirectTarget = "/index"
	} else {
		redirectTarget = "/register"
	}

	http.Redirect(write, request, redirectTarget, 302)

	fmt.Fprintln(write, bool(bool(loginCheck)))
}
