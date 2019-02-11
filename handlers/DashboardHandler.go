package handlers

import (
	"Site1/helpers"
	"fmt"
	"net/http"
)

func DashboardHandler(write http.ResponseWriter, request *http.Request) {
	checkUser := helpers.GetUserCookie(request)

	if len(checkUser) == 0 {
		redirectTarget := "/login"
		http.Redirect(write, request, redirectTarget, 302)
	}

	fmt.Fprintln(write, "Welcome! Joe")
}
