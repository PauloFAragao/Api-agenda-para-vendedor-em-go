package models

import "time"

// Sales representa as vendas feitas ao cliente pelo vendedor
type Sales struct {
	ID       uint64    `json:"id,omitempty"`       // ID no banco de dados - esse dado não vai ser visível para o usuário
	SellerID uint64    `json:"sellerId,omitempty"` // ID do vendedor vinculado a interação - esse dado não vai ser visível para o usuário
	ClientID uint64    `json:"clientId,omitempty"` // ID do cliente vinculado a interação - esse dado não vai ser visível para o usuário
	Date     time.Time `json:"date,omitempty"`     // data da venda
	Sale     string    `json:"sale,omitempty"`     // o que foi vendido
}
