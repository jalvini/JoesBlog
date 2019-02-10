package handlers

import (
	"fmt"
	"net/http"
)

func CommentHandler(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(write, "Welcome! Joe")
}
