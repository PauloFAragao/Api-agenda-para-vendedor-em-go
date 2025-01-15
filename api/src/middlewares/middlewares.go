package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate verifica se o usuário fazendo a requisição está autenticado
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// validando o token
		if err := authentication.ValidateToken(r); err != nil {
			responses.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		nextFunction(w, r)
	}
}
