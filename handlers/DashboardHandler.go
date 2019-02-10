package handlers

import (
	"fmt"
	"net/http"
)

func DashboardHandler(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(write, "Welcome! Joe")
}
