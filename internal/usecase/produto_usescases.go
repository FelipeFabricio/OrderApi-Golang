package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type ProdutoUseCases struct {
	ProdutoDb interfaces.ProdutoDbInterface
}

func NewProdutoUseCases(produtoDb interfaces.ProdutoDbInterface) *ProdutoUseCases {
	return &ProdutoUseCases{ProdutoDb: produtoDb}
}

func (p *ProdutoUseCases) Inserir(produto *entity.Produto) error {
	novoProduto, err := produto.NewProduto(produto.Nome, produto.Descricao, produto.Categoria, produto.Valor)
	if err != nil {
		return err
	}
	return p.ProdutoDb.Inserir(novoProduto)
}

func (p *ProdutoUseCases) ObterTodosProdutos() (*[]entity.Produto, error) {
	return p.ProdutoDb.ObterTodos()
}

func (p *ProdutoUseCases) ObterPorCategoria(categoria entity.CategoriaProduto) (*[]entity.Produto, error) {
	return p.ProdutoDb.ObterPorCategoria(categoria)
}
