package models

import (
	"Site1/database"
)

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func CreatePost() string {
	db := database.DbInit()

	insert, err := db.Query("INSERT INTO posts VALUES ( 3, 'TEST 3', 'Tester 3', 'This is test 3' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	defer db.Close()

	return "Row Updated Successfully"
}

func ReadPost() []Post {
	post := Post{}
	res := []Post{}

	db := database.DbInit()

	defer db.Close()

	results, err := db.Query("SELECT id, title, description, body FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&post.ID, &post.Title, &post.Description, &post.Body)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, post)
	}

	defer db.Close()

	return res
}

func UpdatePost() {

}

func DeletePost() {

}

func ReadPosts() []Post {
	post := Post{}
	res := []Post{}

	db := database.DbInit()

	defer db.Close()

	results, err := db.Query("SELECT id, title, description, body FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&post.ID, &post.Title, &post.Description, &post.Body)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, post)
	}

	defer db.Close()

	return res
}
