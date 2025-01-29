package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s, %s, %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Authenticate verifica a existência dos cookies
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound) // 302
			return
		}

		nextFunc(w, r)
	}
}
