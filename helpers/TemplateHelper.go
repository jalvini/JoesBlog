package helpers

import (
	"html/template"
)

var Template *template.Template

func ParseTemplates() {
	Template = template.Must(template.ParseFiles("views/user/register.gohtml", "views/user/login.gohtml", "views/index.gohtml", "views/shared/header.gohtml", "views/shared/footer.gohtml"))
}
