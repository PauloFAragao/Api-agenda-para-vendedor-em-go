package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse retorna uma resposta em JSON para a requisição
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {

	// para enviar em json
	w.Header().Set("Content-Type", "application/json")

	// escrevendo o http status code
	w.WriteHeader(statusCode)

	if data != nil {
		// colocando os dados em um json
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// ErrorResponse retorna um erro em formato JSON para a requisição
func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	JSONResponse(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}
