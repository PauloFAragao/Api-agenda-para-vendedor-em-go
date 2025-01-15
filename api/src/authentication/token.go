package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateToken retorna um token assinado com as permissões do usuário
func CreateToken(userID uint64) (string, error) {

	// definindo as permissões
	permissions := jwt.MapClaims{}

	// campo autorizado
	permissions["authorized"] = true

	// definindo que o token deve expirar em 6 horas
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()

	// inserindo o id do usuário no token
	permissions["userId"] = userID

	// gerando chave secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	// assinando o token e retornando ele
	return token.SignedString([]byte(config.SecretKey)) // secret
}

// ValidateToken verifica se o token passando na requisição é valido
func ValidateToken(r *http.Request) error {

	// pegando o token
	tokenString := extractToken(r)

	// dando o parse no token
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}

	// validando
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// ExtractUserID retorna o usuárioId que está salvo no token
func ExtractUserID(r *http.Request) (uint64, error) {
	// pegando o token
	tokenString := extractToken(r)

	// dando o parse no token
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return 0, err
	}

	// pegando as permissões
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// pegando o id do usuário
		usuarioID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, nil
		}

		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	// pegando o token
	token := r.Header.Get("Authorization")

	// verificando se a string do token tem 2 palavras (ele tem que vir com a palavra bearer antes do token)
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""

}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	// verificando se o método de assinatura é o esperado
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
