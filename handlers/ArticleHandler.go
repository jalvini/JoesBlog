package handlers

import (
	"Site1/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func ArticleHandler(write http.ResponseWriter, request *http.Request) {
	post := models.ReadPosts()

	res2B, _ := json.Marshal(post)
	fmt.Fprintln(write, string(string(res2B)))
}

func PostArticleHandler(write http.ResponseWriter, request *http.Request) {
	post := models.CreatePost()

	res2B, _ := json.Marshal(post)
	fmt.Fprintln(write, string(string(res2B)))
}

/*
func SingleArticleHandler(write http.ResponseWriter, request *http.Request){
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
}

*/
