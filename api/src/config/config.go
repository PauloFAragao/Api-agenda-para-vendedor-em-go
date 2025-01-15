package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DataBaseConnectionString contém a string para se conectar ao banco de dados
	DataBaseConnectionString = ""

	// ConnectionPort é a porta onde a api vai estar rodando
	ConnectionPort = 0

	// SecretKey chave de segurança para assinar os tokens
	SecretKey []byte
)

// LoadAmbientVariables vai carregar as variáveis de ambiente
func LoadAmbientVariables() {
	var err error

	// lendo do arquivo .env
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// convertendo o valor capturado do .env e atribuindo
	ConnectionPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ConnectionPort = 9000 // porta padrão
	}

	// criando a string de conexão com o banco de dados
	DataBaseConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
