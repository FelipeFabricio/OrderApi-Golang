package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type ClienteUseCases struct {
	ClienteDb interfaces.ClienteDbInterface
}

func NewClienteUseCases(clienteDb interfaces.ClienteDbInterface) *ClienteUseCases {
	return &ClienteUseCases{ClienteDb: clienteDb}
}

func (c *ClienteUseCases) ObterTodos() (*[]entity.Cliente, error) {
	return c.ClienteDb.ObterTodos()
}

func (c *ClienteUseCases) Inserir(cliente *entity.Cliente) error {
	novoCliente, err := cliente.NewCliente(cliente.Nome, cliente.Email, cliente.Cpf)
	if err != nil {
		return err
	}
	return c.ClienteDb.Inserir(novoCliente)
}

func (c *ClienteUseCases) Atualizar(cliente *entity.Cliente) error {
	return c.ClienteDb.Atualizar(cliente)
}
