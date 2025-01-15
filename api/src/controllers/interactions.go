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

// CreateInteraction é chamado quando a rota /interacao com o método post é acessada - é chamado para criar uma interação entre um usuário e um cliente
func CreateInteraction(w http.ResponseWriter, r *http.Request) {
	// capturando o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var interaction models.Interactions

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &interaction); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados
	if err = interaction.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// inserindo o id do usuário na interação
	interaction.SellerID = userIDInToken

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando o repositório
	repository := repository.NewInteractionsRepository(db)

	// mandando criar a interação no banco de dados
	interaction.ID, err = repository.CreateInteractions(interaction)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// enviando a resposta
	responses.JSONResponse(w, http.StatusCreated, interaction)

}

// ViewInteractions é chamado quando a rota /interacao com o método get é acessada - é chamada para pesquisar todas as interações que o usuário teve com seus clientes
func ViewInteractions(w http.ResponseWriter, r *http.Request) {
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
	repository := repository.NewInteractionsRepository(db)

	// buscando
	sales, erro := repository.SearchAllInteractions(userIDInToken)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, sales)
}

// SearchInteraction é chamado quando a rota /interacao/{interactionId} com o método get é acessada - é chamada para pesquisar uma interação especifica, usando seu id
func SearchInteraction(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o parâmetro para inteiro
	interactionID, erro := strconv.ParseUint(parameters["interactionId"], 10, 64)
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
	repository := repository.NewInteractionsRepository(db)

	// buscando
	interactions, erro := repository.SearchByID(userIDInToken, interactionID)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, interactions)
}

// SearchInteractionWithClient é chamado quando a rota /interacao-cliente/{clientId} com o método get é acessada - é chamada para ver todas as interações do usuário com um cliente especifico, usando o id do cliente
func SearchInteractionWithClient(w http.ResponseWriter, r *http.Request) {
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
	repository := repository.NewInteractionsRepository(db)

	// buscando
	clients, erro := repository.SearchByClient(userIDInToken, clientID)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, clients)
}

// EditInteraction é chamado quando a rota /interacao/{interactionId} com o método put é acessada
func EditInteraction(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id da interação em uint64
	interactionID, err := strconv.ParseUint(parameters["interactionId"], 10, 64)
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
	repository := repository.NewInteractionsRepository(db)

	// pegando o id do usuário vinculado a interação
	linkedUserId, err := repository.GetLinkedUserId(interactionID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// verificando se o id do que vai ser alterado e o id no token são iguais
	if linkedUserId != userIDInToken {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível atualizar uma interação que não seja sua"))
		return
	}

	// capturando o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var interaction models.Interactions

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &interaction); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados da interação
	if err = interaction.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// colocando o id da interação na interação
	interaction.ID = interactionID

	// inserindo no banco de dados
	if err := repository.EditInteraction(interaction); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)
}

// DeleteInteraction é chamado quando a rota /interacao/{interactionId} com o método delete é acessada
func DeleteInteraction(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id da interação em uint64
	interactionID, err := strconv.ParseUint(parameters["interactionId"], 10, 64)
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
	repository := repository.NewInteractionsRepository(db)

	// pegando o id do usuário vinculado a interação
	linkedUserId, err := repository.GetLinkedUserId(interactionID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// verificando se o id do que vai ser alterado e o id no token são iguais
	if linkedUserId != userIDInToken {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível deletar uma interação que não seja sua"))
		return
	}

	// inserindo no banco de dados
	if err := repository.DisableInteraction(interactionID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)
}
