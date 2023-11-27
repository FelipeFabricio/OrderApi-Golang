package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrDadosInvalidados = errors.New("Dados do Cliente inválidos")
)

type Cliente struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
}

func (c *Cliente) NewCliente(nome, email, cpf string) (*Cliente, error) {

	if !c.ValidarDadosClientes(nome, email, cpf) {
		return nil, ErrDadosInvalidados
	}

	novoCliente := &Cliente{
		ID:    uuid.New(),
		Nome:  nome,
		Cpf:   cpf,
		Email: email,
	}
	return novoCliente, nil
}

func (c *Cliente) ValidarDadosClientes(nome, email, cpf string) bool {
	return nome != "" && email != "" && cpf != ""
}
