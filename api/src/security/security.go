package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword comprara uma senha e um hash e retorna se eles são iguais
func VerifyPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
