package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(write http.ResponseWriter, request *http.Request) {
	template, _ := template.ParseFiles("views/index.html", "views/shared/header.html", "views/shared/footer.html")
	template.Execute(write, "This is a template")
}
