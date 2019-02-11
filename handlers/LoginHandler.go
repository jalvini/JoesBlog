package handlers

import (
	"Site1/helpers"
	"Site1/models"
	"net/http"
)

func LoginHandler(write http.ResponseWriter, request *http.Request) {
	checkUser := helpers.GetUserCookie(request)
	message := models.Message{}

	if len(checkUser) > 0 {
		redirectTarget := "/dashboard"
		http.Redirect(write, request, redirectTarget, 302)
	}

	write.Header().Set("Content-Type", "text/html")

	err := helpers.Template.ExecuteTemplate(write, "login.gohtml", message)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPostHandler(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	username := request.Form.Get("username")
	password := request.Form.Get("password")
	message := models.Message{"Login Credentials Are Incorrect Please Re-enter Your Username And Password"}

	loginCheck := models.UserGuard(username, password)

	if loginCheck == true {
		helpers.SetUserCookie(username, write)
		redirectTarget := "/dashboard"
		http.Redirect(write, request, redirectTarget, 302)

	}

	write.Header().Set("Content-Type", "text/html")

	err := helpers.Template.ExecuteTemplate(write, "login.gohtml", message)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}
