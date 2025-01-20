package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse representa a resposta de erro da api
type ErrorAPI struct {
	ErrAPI string `json:"erro"`
}

// JSONResponse retorna uma resposta em JSON para a requisição
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {

	// para enviar em json
	w.Header().Set("Content-Type", "application/json")

	// escrevendo o http status code
	w.WriteHeader(statusCode)

	//if data != nil {
	// colocando os dados em um json
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
	//}
}

// HandleErrorStatusCode trata as requisições com status code 400 ou superior
func HandleErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI

	json.NewDecoder(r.Body).Decode(&err)

	JSONResponse(w, r.StatusCode, err)

}
