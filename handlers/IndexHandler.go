package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "text/html")

	template, err := template.ParseFiles("views/index.gohtml", "views/shared/header.gohtml", "views/shared/footer.gohtml")

	if err != nil {
		fmt.Println(write, "Failed To Load Template")
		panic(err.Error())
	}

	tempErr := template.Execute(write, "John Smith")
	if tempErr != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}
