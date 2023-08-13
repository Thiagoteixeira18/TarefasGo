package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template 

func CarregarTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplete(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}