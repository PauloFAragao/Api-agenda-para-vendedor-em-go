package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// PerformLogin utiliza o e-mail e senha o usuário para autenticar na aplicação
func PerformLogin(w http.ResponseWriter, r *http.Request) {
	// vai pegar o corpo da requisição
	r.ParseForm()

	// map para criar o json para enviar para a api
	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSONResponse(w, http.StatusBadRequest, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}

	// url pra enviar para api
	url := fmt.Sprintf("%s/login", config.APIURL)

	// enviando para a api
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSONResponse(w, http.StatusInternalServerError, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o statusCode resposta da api
	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var authenticationData models.AuthenticationData

	// pegando os dados do json
	if err = json.NewDecoder(response.Body).Decode(&authenticationData); err != nil {
		responses.JSONResponse(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}

	// criando o cookie
	if err = cookies.Save(w, authenticationData.ID, authenticationData.Token); err != nil {
		responses.JSONResponse(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)

}
