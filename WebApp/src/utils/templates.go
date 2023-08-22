package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template 

func CarregarTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

func ExecutarTemplete(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}