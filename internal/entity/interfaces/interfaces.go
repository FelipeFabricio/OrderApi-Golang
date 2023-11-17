package database

import "github.com/felipefabricio/wonder-food/internal/entity"

type ProdutoInterface interface {
	Inserir(produto *entity.Produto) error
	ObterTodos(pagina, limite int, sort string) (*[]entity.Produto, error)
	ObterPorCategoria(categoria entity.CategoriaProduto) (*entity.Produto, error)
	Atualizar(produto *entity.Produto) error
	Deletar(id string) error
}

type ClienteInterface interface {
	Inserir(cliente *entity.Cliente) error
	Atualizar(cliente *entity.Cliente) error
}

type PedidoInterface interface {
	Inserir(pedido *entity.Pedido, produtos *[]entity.Produto) (*entity.Pedido, error)
	ObterTodos(pagina, limite int, sort string) (*[]entity.Pedido, error)
	ObterPorNumeroPedido(numeroPedido string) (*entity.Pedido, error)
	Atualizar(pedido *entity.Pedido) error
	Deletar(id string) error
}
