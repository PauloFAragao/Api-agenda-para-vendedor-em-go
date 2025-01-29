package models

import "time"

// Interactions representa as interações feitas entre vendedor e cliente
type Interactions struct {
	ClientName  string    `json:"clientName,omitempty"`
	Status      string    `json:"status,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Interaction string    `json:"interaction,omitempty"`
	Content     string    `json:"content,omitempty"`
}
