package models

import (
	"errors"
	"time"
)

// Sales representa as vendas feitas ao cliente pelo vendedor
type Sales struct {
	ID       uint64    `json:"id,omitempty"`       // ID no banco de dados - esse dado não vai ser visível para o usuário
	SellerID uint64    `json:"sellerId,omitempty"` // ID do vendedor vinculado a interação - esse dado não vai ser visível para o usuário
	ClientID uint64    `json:"clientId,omitempty"` // ID do cliente vinculado a interação - esse dado não vai ser visível para o usuário
	Date     time.Time `json:"date,omitempty"`     // data da venda
	Sale     string    `json:"sale,omitempty"`     // o que foi vendido
	Active   bool      `json:"active,omitempty"`   // indica se a venda está ativo ou não - esse dado não vai ser visível para o usuário
}

// Sale existe para ser usado como resposta
type Sale struct {
	Date time.Time `json:"date,omitempty"` // data da venda
	Sale string    `json:"sale,omitempty"` // o que foi vendido
}

// Prepare vai chamar os métodos para validar os dados da venda
func (sale *Sales) Prepare() error {

	// validando os dados
	if err := sale.validate(); err != nil {
		return err
	}

	return nil
}

func (sale *Sales) validate() error {

	// verificando se tem id
	if sale.ClientID <= 0 {
		return errors.New("é necessário ter o id do cliente")
	}

	// verificar Date
	if ok := sale.Date.IsZero(); ok {
		return errors.New("a data não pode estar em branco")
	}

	// verificar Interaction
	if sale.Sale == "" {
		return errors.New("é necessário informar a venda")
	}

	return nil
}
