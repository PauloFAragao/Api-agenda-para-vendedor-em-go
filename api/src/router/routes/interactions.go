package routes

import (
	"api/src/controllers"
	"net/http"
)

var interactionRotes = []Route{
	{ // rota para criar uma nova interação
		URI:                "/interacao",
		Method:             http.MethodPost,
		Function:           controllers.CreateInteraction,
		NeedAuthentication: false,
	},
	{ // rota para ver todas as interações de um usuário
		URI:                "/interacao",
		Method:             http.MethodGet,
		Function:           controllers.ViewInteractions,
		NeedAuthentication: false,
	},
	{ // rota para ver uma interação especifica entre o usuário e um cliente
		URI:                "/interacao/{interactionId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchInteraction,
		NeedAuthentication: false,
	},
	{ // rota para ver todas as interações entre o usuário e um cliente
		URI:                "/interacao-cliente/{clientId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchInteractionWithClient,
		NeedAuthentication: false,
	},
	{ // rota para ver todas as interações marcadas
		URI:                "/interacoes-marcadas",
		Method:             http.MethodGet,
		Function:           controllers.ViewTaggedInteractions,
		NeedAuthentication: true,
	},
	{ // rota para editar uma interação
		URI:                "/interacao/{interactionId}",
		Method:             http.MethodPut,
		Function:           controllers.EditInteraction,
		NeedAuthentication: false,
	},
	{ // rota para deletar uma interação
		URI:                "/interacao/{interactionId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteInteraction,
		NeedAuthentication: false,
	},
	{ // rota para buscar todas as interações de um dia
		URI:                "/interacao-date/{interactionDate}",
		Method:             http.MethodGet,
		Function:           controllers.ViewInteractionsMarkedOnDate,
		NeedAuthentication: false,
	},
}
