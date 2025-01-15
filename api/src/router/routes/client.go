package routes

import (
	"api/src/controllers"
	"net/http"
)

var clientRoutes = []Route{
	{ // rota para criar um novo cliente
		URI:                "/cliente",
		Method:             http.MethodPost,
		Function:           controllers.CreateClient,
		NeedAuthentication: true,
	},
	{ // rota para visualizar todos os clientes do usuário
		URI:                "/cliente",
		Method:             http.MethodGet,
		Function:           controllers.ViewClients,
		NeedAuthentication: true,
	},
	{ // rota para visualizar um cliente por id
		URI:                "/cliente/{clientId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchClient,
		NeedAuthentication: true,
	},
	{ // rota para editar um cliente
		URI:                "/cliente/{clientId}",
		Method:             http.MethodPut,
		Function:           controllers.EditClient,
		NeedAuthentication: true,
	},
	{ // rota para deletar um cliente
		URI:                "/cliente/{clientId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteClient,
		NeedAuthentication: true,
	},
	{ // rota para buscar um usuário
		URI:                "/buscar-cliente",
		Method:             http.MethodGet,
		Function:           controllers.SearchClientByName,
		NeedAuthentication: true,
	},
}
