package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da api
type Route struct {
	URI                string                                   // caminho
	Method             string                                   // GET, POST, PUT, DELETE...
	Function           func(http.ResponseWriter, *http.Request) // Função a ser chamada
	NeedAuthentication bool                                     // se essa pagina precisa de autenticação
}

// Configurar coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {

	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, clientRoutes...)
	routes = append(routes, interactionRotes...)
	routes = append(routes, salesRoutes...)

	// adicionando as rotas ao router
	for _, route := range routes {

		// verificando se a rota precisa de autenticação
		if route.NeedAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
