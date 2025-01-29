package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

// LoadLoginScreen vai renderizar a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadHomeScreen vai carrega a pagina home
func LoadHomeScreen(w http.ResponseWriter, r *http.Request) {

	// url para conectar com a api
	url := fmt.Sprintf("%s/interacao-date/%s", config.APIURL, time.Now().Format("2006-01-02"))

	// resposta da api
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("erro na resposta")
		responses.JSONResponse(w, http.StatusInternalServerError, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code do html
	if response.StatusCode >= 400 {
		fmt.Println("erro no status code: ", response.StatusCode)
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var interactions []models.Interactions

	// colocando as publicações no json
	if err = json.NewDecoder(response.Body).Decode(&interactions); err != nil {
		fmt.Println("erro no JSON")
		responses.JSONResponse(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrAPI: err.Error()})
		return
	}

	// pegando o cookie para depois pegar o id do usuário
	cookie, _ := cookies.Read(r)

	// pegando o id
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// renderizando a pagina de login
	utils.ExecuteTemplate(w, "home.html", struct {
		Interactions []models.Interactions
		UserID       uint64
	}{
		Interactions: interactions,
		UserID:       userID,
	})

}
