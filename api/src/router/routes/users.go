package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{ // rota para criar um novo usuário
		URI:                "/usuario",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},
	{ // rota para editar um usuário
		URI:                "/usuario/{userId}",
		Method:             http.MethodPut,
		Function:           controllers.EditUser,
		NeedAuthentication: true,
	},
	{ // rota para deletar um usuário
		URI:                "/usuario/{userId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},
}
