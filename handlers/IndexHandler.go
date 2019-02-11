package handlers

import (
	"Site1/helpers"
	"Site1/models"
	"fmt"
	"net/http"
)

func IndexHandler(write http.ResponseWriter, request *http.Request) {
	user := helpers.GetUserCookie(request)
	userData := models.User{}

	if len(user) > 0 {
		userData = models.ShowUser(user)
		fmt.Println("This Works")
	}

	fmt.Println(userData)

	write.Header().Set("Content-Type", "text/html")

	err := helpers.Template.ExecuteTemplate(write, "index.gohtml", "Joe")
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}
