package models

// AuthenticationData contém o token e o id do usuário autenticado
type AuthenticationData struct {
	ID    string `json: "id"`
	Token string `json: "token"`
}
