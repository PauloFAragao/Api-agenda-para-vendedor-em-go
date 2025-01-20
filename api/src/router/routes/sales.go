package routes

import (
	"api/src/controllers"
	"net/http"
)

var salesRoutes = []Route{
	{ // criar uma venda
		URI:                "/venda",
		Method:             http.MethodPost,
		Function:           controllers.CreateSale,
		NeedAuthentication: false,
	},
	{ // visualizar todas as vendas
		URI:                "/venda",
		Method:             http.MethodGet,
		Function:           controllers.ViewSales,
		NeedAuthentication: false,
	},
	{ // visualizar uma venda
		URI:                "/venda/{sellId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchSale,
		NeedAuthentication: false,
	},
	{ // visualizar todas as vendas feitas a um cliente
		URI:                "/vendas-cliente/{clientId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchSalesForClient,
		NeedAuthentication: false,
	},
	{ // editar uma venda
		URI:                "/venda/{sellId}",
		Method:             http.MethodPut,
		Function:           controllers.EditSale,
		NeedAuthentication: false,
	},
	{ // deletar uma venda
		URI:                "/venda/{sellId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteSale,
		NeedAuthentication: false,
	},
}
