package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoute = Route{
	URI:                "/usuarios",
	Method:             http.MethodPost,
	Function:           controllers.CreateUser,
	NeedAuthentication: false,
}
