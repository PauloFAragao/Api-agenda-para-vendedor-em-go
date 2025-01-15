package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser é chamado quando a rota /usuario com o método post é acessada - insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

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

	// validando os dados do usuário
	if err = user.Prepare(0); err != nil {
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

	// mandando criar o usuário no banco de dados
	user.ID, err = repository.CreateUser(user)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// enviando a resposta
	responses.JSONResponse(w, http.StatusCreated, user)

}

// EditUser é chamado quando a rota /usuario/{userId} com o método put é acessada - edita um usuário no banco de dados
func EditUser(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id do usuário em uint64
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// verificando se o id que vai ser alterado e o id no token são iguais
	if userID != userIDInToken {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

	// lendo o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	// extraindo dados do json
	if err = json.Unmarshal(requisitionBody, &user); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados do usuário
	if err = user.Prepare(1); err != nil {
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

	// mandando editar o usuário
	if err = repository.EditUser(userID, user); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)

}

// Delete é chamado quando a rota /usuario/{userId} com o método delete é acessada - user faz um soft delete do usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id do usuário em uint64
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// verificando se o id que vai ser alterado e o id no token são iguais
	if userID != userIDInToken {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível deletar um usuário que não seja o seu"))
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

	// mandando desativar o usuário
	if err = repository.DisableUser(userID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}
