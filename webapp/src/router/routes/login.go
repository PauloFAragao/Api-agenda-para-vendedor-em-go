package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:                "/",
		Method:             http.MethodGet,
		Function:           controllers.LoadLoginScreen,
		NeedAuthentication: false,
	},
	{
		URI:                "/login",
		Method:             http.MethodGet,
		Function:           controllers.LoadLoginScreen,
		NeedAuthentication: false,
	},
	{
		URI:                "/login",
		Method:             http.MethodPost,
		Function:           controllers.PerformLogin,
		NeedAuthentication: false,
	},
}
