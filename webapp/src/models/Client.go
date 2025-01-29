package models

// Clients representa um cliente que foi a
type Clients struct {
	Name     string `json:"name,omitempty"`
	Contacts string `json:"contacts,omitempty"` // telefone, WhatsApp, etc...
	Address  string `json:"address,omitempty"`  // endereço completo do cliente
}
