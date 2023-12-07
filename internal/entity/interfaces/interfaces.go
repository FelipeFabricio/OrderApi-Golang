package entity

import (
	"github.com/felipefabricio/wonder-food/internal/dto"
	"github.com/felipefabricio/wonder-food/internal/entity"
)

type ProdutoUseCasesInterface interface {
	Inserir(produto *dto.CriarProdutoInputDto) error
	ObterTodosProdutos() (*[]dto.ObterProdutoOutputDto, error)
	ObterPorCategoria(categoria entity.CategoriaProduto) (*[]dto.ObterProdutoOutputDto, error)
}
type ProdutoDbInterface interface {
	Inserir(produto *entity.Produto) error
	ObterTodos() (*[]entity.Produto, error)
	ObterPorCategoria(categoria entity.CategoriaProduto) (*[]entity.Produto, error)
	ObterPorId(id string) (*entity.Produto, error)
	Atualizar(produto *entity.Produto) error
	Deletar(id string) error
}

type ClienteUseCasesInterface interface {
	Inserir(cliente *dto.CriarClienteInputDto) error
	Atualizar(cliente *dto.AtualizarClienteInputDto) error
	ObterTodos() (*[]dto.ObterClienteOutputDto, error)
}
type ClienteDbInterface interface {
	Inserir(cliente *entity.Cliente) error
	Atualizar(cliente *entity.Cliente) error
	ObterTodos() (*[]entity.Cliente, error)
}

type PedidoUseCasesInterface interface {
	ObterPedidosEmAberto() (*[]dto.ObterPedidosOutputDto, error)
	Inserir(pedido *dto.CriarPedidoInputDto) error
	ConsultarStatusPagamento(numeroPedido int) (*dto.ObterStatusPedidoOutputDto, error)
}
type PedidoDbInterface interface {
	ObterPedidosEmAberto() (*[]entity.Pedido, error)
	Inserir(pedido *entity.Pedido) error
	ObterPorNumeroPedido(numeroPedido int) (*entity.Pedido, error)
	AtualizarStatusPagamento(numeroPedido int, status entity.StatusPedido) error
	DeletarPedido(numeroPedido int) error
}
