package dto

import (
	"github.com/google/uuid"
)

// TODO: Criar Dtos específicas para cada Caso de uso (inputs e outputs)
type ClienteDto struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
}
