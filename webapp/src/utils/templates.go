package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates insere os templates html na variável templates
func LoadTemplates() {
	templates = template.Must(templates.ParseGlob("views/*.html"))
}

// ExecuteTemplate renderiza uma pagina html na tela
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
