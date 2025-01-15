package models

import (
	"errors"
	"time"
)

// Interactions representa as interações feitas entre vendedor e cliente
type Interactions struct {
	ID          uint64    `json:"id,omitempty"`          // ID no banco de dados - esse dado não vai ser visível para o usuário
	SellerID    uint64    `json:"sellerId,omitempty"`    // ID do vendedor vinculado a interação - esse dado não vai ser visível para o usuário
	ClientID    uint64    `json:"clientId,omitempty"`    // ID do cliente vinculado a interação - esse dado não vai ser visível para o usuário
	Status      string    `json:"status,omitempty"`      // se a interação está marcada ou se já aconteceu - talvez tenha que mudar o dado para int
	Date        time.Time `json:"date,omitempty"`        // data da interação
	Interaction string    `json:"interaction,omitempty"` // tipo de interação: pessoalmente, por telefone, por mensagem...
	Content     string    `json:"content,omitempty"`     // o que foi conversado
	Active      bool      `json:"active,omitempty"`      // se está ativo
}

// Prepare vai chamar os métodos para validar os dados da interação
func (interaction *Interactions) Prepare() error {

	// validando os dados
	if err := interaction.validate(); err != nil {
		return err
	}

	return nil
}

func (interaction *Interactions) validate() error {

	// verificando se tem id
	if interaction.ClientID <= 0 {
		return errors.New("é necessário ter o id do cliente")
	}

	// verificando se tem dados de status
	if interaction.Status == "" {
		return errors.New("o status é obrigatório e não pode estar em branco, complete com: já aconteceu, marcada, re-marcada, etc")
	}

	// verificar Date
	if ok := interaction.Date.IsZero(); ok {
		return errors.New("a data não pode estar em branco")
	}

	// verificar Interaction
	if interaction.Interaction == "" {
		return errors.New("o tipo de interação não pode estar em branco, complete com: pessoalmente, por telefone, por mensagem, etc")
	}

	return nil
}
