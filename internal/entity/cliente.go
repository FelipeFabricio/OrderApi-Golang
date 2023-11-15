package entity

import "github.com/google/uuid"

type Cpf struct {
	Cpf string `json:"cpf"`
}

type Cliente struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Cpf   Cpf       `json:"cpf"`
	Email string    `json:"email"`
}
