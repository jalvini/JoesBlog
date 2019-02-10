package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(":8000", r))
}
