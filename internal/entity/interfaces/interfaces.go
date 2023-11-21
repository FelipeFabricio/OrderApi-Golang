package entity

import "github.com/felipefabricio/wonder-food/internal/entity"

type ProdutoRepositoryInterface interface {
	Inserir(produto *entity.Produto) error
	ObterTodos() (*[]entity.Produto, error)
	ObterPorCategoria(categoria entity.CategoriaProduto) (*[]entity.Produto, error)
	Atualizar(produto *entity.Produto) error
	Deletar(id string) error
}

type ClienteRepositoryInterface interface {
	Inserir(cliente *entity.Cliente) error
	Atualizar(cliente *entity.Cliente) error
}

type PedidoRepositoryInterface interface {
	Inserir(pedido *entity.Pedido, produtos *[]entity.Produto) (*entity.Pedido, error)
	ObterTodos(pagina, limite int, sort string) (*[]entity.Pedido, error)
	ObterPorNumeroPedido(numeroPedido string) (*entity.Pedido, error)
	Atualizar(pedido *entity.Pedido) error
	Deletar(id string) error
}

type ProdutoUseCasesInterface interface {
	Inserir(produto *entity.Produto) error
	ObterTodosProdutos() (*[]entity.Produto, error)
	ObterPorCategoria(categoria entity.CategoriaProduto) (*[]entity.Produto, error)
}
