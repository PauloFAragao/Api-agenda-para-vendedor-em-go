package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Login é responsável por autenticar um usuário na api
func Login(w http.ResponseWriter, r *http.Request) {

	// capturando o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &user); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando o repositório
	repository := repository.NewUsersRepository(db)

	// pegando os dados do usuário que estão no banco
	userSavedInDatabase, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// verificando se a senha que o usuário enviou está correta
	if err = security.VerifyPassword(userSavedInDatabase.Password, user.Password); err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// gerando o token
	token, err := authentication.CreateToken(userSavedInDatabase.ID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userSavedInDatabase.ID, 10)

	responses.JSONResponse(w, http.StatusOK, models.AuthenticationData{ID: userID, Token: token})

	//w.Write([]byte(token))
}
