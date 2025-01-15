package models

import (
	"errors"
	"strings"
)

// Clients representa um cliente que foi a
type Clients struct {
	ID       uint64 `json:"id,omitempty"`       // ID no banco de dados - esse dado não vai ser visível para o usuário
	SellerID uint64 `json:"sellerId,omitempty"` // ID do vendedor vinculado ao cliente - esse dado não vai ser visível para o usuário
	Name     string `json:"name,omitempty"`
	Contacts string `json:"contacts,omitempty"` // telefone, WhatsApp, etc...
	Address  string `json:"address,omitempty"`  // endereço completo do cliente
	Active   bool   `json:"active,omitempty"`   // indica se o usuário está ativo ou não - esse dado não vai ser visível para o usuário
}

// Prepare vai chamar os métodos para formatar e validar os dados do cliente
func (client *Clients) Prepare() error {

	// retirando os espaços em branco
	client.format()

	// validando os dados
	if err := client.validate(); err != nil {
		return err
	}

	return nil
}

func (client *Clients) validate() error {

	// verificando se o nome está em branco
	if client.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	return nil
}

func (client *Clients) format() {
	client.Name = strings.TrimSpace(client.Name)
	client.Address = strings.TrimSpace(client.Address)
}
