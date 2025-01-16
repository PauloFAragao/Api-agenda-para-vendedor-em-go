package models

// Password representa o formato da requisição de alteração de senha
type Password struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
