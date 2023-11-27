package dto

import (
	"github.com/google/uuid"
)

type CriarClienteInputDto struct {
	Nome  string `json:"nome"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
}

type ObterClienteOutputDto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
}

type AtualizarClienteInputDto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
}
