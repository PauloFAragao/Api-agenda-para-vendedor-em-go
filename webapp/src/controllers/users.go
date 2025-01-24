package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/responses"
)

// CreateUser chama a api para cadastrar um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// vai pegar o corpo da requisição
	r.ParseForm()

	// map para criar o json para enviar para a api
	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSONResponse(w, http.StatusBadRequest, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}

	// url pra enviar para api
	url := fmt.Sprintf("%s/usuario", config.APIURL)

	// enviando para a api
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSONResponse(w, http.StatusInternalServerError, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code
	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	// enviando a resposta para o usuário
	responses.JSONResponse(w, response.StatusCode, nil)

}
