package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginScreen vai renderizar a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
