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

// CreateSale é chamado quando a rota /venda" com o método post é acessada - é chamado para criar uma venda
func CreateSale(w http.ResponseWriter, r *http.Request) {

	// capturando o corpo da requisição
	requisitionBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var sale models.Sales

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &sale); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados
	if err = sale.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// pegando o id do usuário no token
	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// inserindo o id do usuário
	sale.SellerID = userIDInToken

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando o repositório
	repository := repository.NewSalesRepository(db)

	// mandando criar a interação no banco de dados
	sale.ID, err = repository.CreateSale(sale)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// enviando a resposta
	responses.JSONResponse(w, http.StatusCreated, sale)
}

// ViewSales é chamado quando a rota /venda com o método get é acessada - é chamado para ver todas as vendas feitas
func ViewSales(w http.ResponseWriter, r *http.Request) {
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
	repository := repository.NewSalesRepository(db)

	// buscando
	clients, erro := repository.SearchAllSales(userIDInToken)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, clients)
}

// SearchSale é chamado quando a rota /venda/{sellId} com o método get é acessada - é chamado para ver uma venda especifica
func SearchSale(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o parâmetro para inteiro
	SaleID, erro := strconv.ParseUint(parameters["sellId"], 10, 64)
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
	repository := repository.NewSalesRepository(db)

	// buscando
	sales, erro := repository.SearchByID(userIDInToken, SaleID)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, sales)
}

// SearchSalesForClient é chamado quando a rota /vendas-cliente/{clientId} com o método get é acessada - é chamado para ver todas as vendas feitas a um cliente
func SearchSalesForClient(w http.ResponseWriter, r *http.Request) {
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
	repository := repository.NewSalesRepository(db)

	// buscando
	clients, erro := repository.SearchByClient(userIDInToken, clientID)
	if erro != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSONResponse(w, http.StatusOK, clients)
}

// EditSale é chamado quando a rota /venda/{sellId} com o método put é acessada
func EditSale(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id da interação em uint64
	saleID, err := strconv.ParseUint(parameters["sellId"], 10, 64)
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
	repository := repository.NewSalesRepository(db)

	// pegando o id do usuário vinculado a interação
	linkedUserId, err := repository.GetLinkedUserId(saleID)
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

	var sale models.Sales

	// lendo o json
	if err = json.Unmarshal(requisitionBody, &sale); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// validando os dados da interação
	if err = sale.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// colocando o id da interação na interação
	sale.ID = saleID

	// inserindo no banco de dados
	if err := repository.EditSale(sale); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)
}

// DeleteSale é chamado quando a rota /venda/{sellId} com o método delete é acessada
func DeleteSale(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parameters := mux.Vars(r)

	// convertendo o id da interação em uint64
	saleID, err := strconv.ParseUint(parameters["sellId"], 10, 64)
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
	repository := repository.NewSalesRepository(db)

	// pegando o id do usuário vinculado a interação
	linkedUserId, err := repository.GetLinkedUserId(saleID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// verificando se o id do que vai ser alterado e o id no token são iguais
	if linkedUserId != userIDInToken {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("não é possível deletar uma venda que não seja sua"))
		return
	}

	// inserindo no banco de dados
	if err := repository.DisableSale(saleID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, nil)
}
