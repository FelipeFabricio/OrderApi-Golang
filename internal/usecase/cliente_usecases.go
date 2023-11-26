package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type ClienteUseCases struct {
	clienteRepository interfaces.ClienteRepositoryInterface
}

func NewClienteUseCases(clienteRepository interfaces.ClienteRepositoryInterface) *ClienteUseCases {
	return &ClienteUseCases{clienteRepository: clienteRepository}
}

func (c *ClienteUseCases) ObterTodos() (*[]entity.Cliente, error) {
	return c.clienteRepository.ObterTodos()
}

func (c *ClienteUseCases) Inserir(cliente *entity.Cliente) error {
	novoCliente, err := cliente.NewCliente(cliente.Nome, cliente.Email, cliente.Cpf)
	if err != nil {
		return err
	}
	return c.clienteRepository.Inserir(novoCliente)
}

func (c *ClienteUseCases) Atualizar(cliente *entity.Cliente) error {
	return c.clienteRepository.Atualizar(cliente)
}
