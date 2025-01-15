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
	"strings"

	"github.com/gorilla/mux"
)

// CreateClient é chamado quando a rota /cliente com o método post é acessada - cria um cliente
func CreateClient(w http.ResponseWriter, r *http.Request) {
	// capturando o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var client models.Clients

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &client); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados do cliente
	if err = client.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// inserindo o id do usuário no cliente
	client.SellerID = userIDInToken

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando o repositório
	repository := repository.NewClientsRepository(db)

	// mandando criar o cliente no banco de dados
	client.ID, err = repository.CreateClient(client)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// enviando a resposta
	responses.JSONResponse(w, http.StatusCreated, client)

}

// ViewClients é chamado quando a rota /cliente com o método get é acessada - busca todos os clientes que o usuário tem cadastrados
func ViewClients(w http.ResponseWriter, r *http.Request) {
	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
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
	repository := repository.NewClientsRepository(db)

	// buscando clientes
	clients, erro := repository.SearchAllClients(userIDInToken)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, clients)
}

// SearchClient é chamado quando a rota /cliente/{clientId} com o método get é acessada - busca um cliente que o usuário tem cadastrado por id
func SearchClient(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o parâmetro para inteiro
	clientID, erro := strconv.ParseUint(parameters["clientId"], 10, 64)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
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
	repository := repository.NewClientsRepository(db)

	// buscando usuário
	client, erro := repository.SearchByID(clientID, userIDInToken)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, client)
}

// EditClient é chamado quando a rota /cliente/{clientId} com o método put é acessada - edita um cliente
func EditClient(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id do client em uint64
	clientID, err := strconv.ParseUint(parameters["clientId"], 10, 64)
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

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando o repositório
	repository := repository.NewClientsRepository(db)

	// pegando o id do usuário que vinculado ao cliente
	linkedUserId, err := repository.GetLinkedUserId(clientID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// verificando se o usuário tentando fazer a edição é o usuário vinculado ao cliente
	if userIDInToken != linkedUserId {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível editar um cliente que não seja seu"))
		return
	}

	// capturando o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var client models.Clients

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &client); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados do cliente
	if err = client.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// colocando o id do cliente na variável
	client.ID = clientID

	// inserindo no banco de dados
	if err := repository.EditClient(client); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)
}

// DeleteClient é chamado quando a rota /cliente/{clientId} com o método delete é acessada - deleta um cliente
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id do client em uint64
	clientID, err := strconv.ParseUint(parameters["clientId"], 10, 64)
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

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando o repositório
	repository := repository.NewClientsRepository(db)

	// pegando o id do usuário que vinculado ao cliente
	linkedUserId, err := repository.GetLinkedUserId(clientID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// verificando se o usuário tentando fazer a deleção é o usuário vinculado ao cliente
	if userIDInToken != linkedUserId {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível deletar um cliente que não seja seu"))
		return
	}

	// fazendo o soft delete
	if err := repository.DisableClient(clientID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)
}

// SearchClientByName é chamado para buscar um cliente usando uma string, quando a rota /cliente/buscar com o método get é acessada - busca um cliente que o usuário tem cadastrado pelo nome
func SearchClientByName(w http.ResponseWriter, r *http.Request) {

	// pegando a string pesquisada
	clientName := strings.ToLower(r.URL.Query().Get("client"))

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
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
	repository := repository.NewClientsRepository(db)

	// buscando usuários
	clients, erro := repository.SearchByName(clientName, userIDInToken)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, clients)

}
