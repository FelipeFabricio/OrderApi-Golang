package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/dto"
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type ClienteUseCases struct {
	ClienteDb interfaces.ClienteDbInterface
}

func NewClienteUseCases(clienteDb interfaces.ClienteDbInterface) *ClienteUseCases {
	return &ClienteUseCases{ClienteDb: clienteDb}
}

func (c *ClienteUseCases) ObterTodos() (*[]dto.ObterClienteOutputDto, error) {
	clientes, err := c.ClienteDb.ObterTodos()
	if err != nil {
		return nil, err
	}

	var clientesDto []dto.ObterClienteOutputDto
	for _, cliente := range *clientes {
		clientesDto = append(clientesDto, dto.ObterClienteOutputDto{
			ID:    cliente.ID,
			Nome:  cliente.Nome,
			Cpf:   cliente.Cpf,
			Email: cliente.Email,
		})
	}
	return &clientesDto, nil
}

func (c *ClienteUseCases) Inserir(cliente *dto.CriarClienteInputDto) error {
	//TODO: Da pra ficar melhor
	var novoCliente entity.Cliente
	clienteValidado, err := novoCliente.NewCliente(cliente.Nome, cliente.Email, cliente.Cpf)
	if err != nil {
		return err
	}
	return c.ClienteDb.Inserir(clienteValidado)
}

func (c *ClienteUseCases) Atualizar(cliente *dto.AtualizarClienteInputDto) error {
	clienteEntity := entity.Cliente{
		ID:    cliente.ID,
		Nome:  cliente.Nome,
		Cpf:   cliente.Cpf,
		Email: cliente.Email,
	}
	return c.ClienteDb.Atualizar(&clienteEntity)
}
