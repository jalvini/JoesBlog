package routes

import (
	"Site1/handlers"
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Router []Route

var Routes = Router{
	Route{
		"IndexHandler",
		"GET",
		"/",
		handlers.IndexHandler,
	},
	Route{
		"ArticleHandler",
		"GET",
		"/article",
		handlers.ArticleHandler,
	},
	Route{
		"PostArticleHandler",
		"POST",
		"/article",
		handlers.PostArticleHandler,
	},
	Route{
		"CommentHandler",
		"GET",
		"/comment",
		handlers.CommentHandler,
	},
	Route{
		"LoginHandler",
		"GET",
		"/login",
		handlers.LoginHandler,
	},
	Route{
		"LoginPostHandler",
		"POST",
		"/login",
		handlers.LoginPostHandler,
	},
	Route{
		"RegisterHandler",
		"GET",
		"/register",
		handlers.RegisterHandler,
	},
	Route{
		"RegisterPostHandler",
		"POST",
		"/register",
		handlers.RegisterPostHandler,
	},
	Route{
		"DashboardHandler",
		"GET",
		"/dashboard",
		handlers.DashboardHandler,
	},
}
