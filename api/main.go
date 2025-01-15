package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadAmbientVariables() // carregando as variáveis de ambiente

	r := router.Generate()

	fmt.Printf("Escutando na porta: %d", config.ConnectionPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ConnectionPort), r))

}
