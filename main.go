package main

import (
	"Site1/helpers"
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	helpers.ParseTemplates()

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(":8000", r))
}
