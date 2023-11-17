package entity

import "github.com/google/uuid"

type ClienteDto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
}
