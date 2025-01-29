package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da aplicação web
type Route struct {
	URI                string                                   // caminho
	Method             string                                   // GET, POST, PUT, DELETE...
	Function           func(http.ResponseWriter, *http.Request) // Função a ser chamada
	NeedAuthentication bool                                     // se essa pagina precisa de autenticação
}

// Configurar coloca todas as rotas dentro do router
func Configure(router *mux.Router) *mux.Router {

	routes := loginRoutes
	routes = append(routes, userRoute...)
	routes = append(routes, homeRoute)

	for _, route := range routes {

		if route.NeedAuthentication {

			router.HandleFunc(
				route.URI, middlewares.Logger(
					middlewares.Authenticate(
						route.Function,
					),
				),
			).Methods(route.Method)

		} else {

			router.HandleFunc(
				route.URI, middlewares.Logger(
					route.Function,
				),
			).Methods(route.Method)

		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
