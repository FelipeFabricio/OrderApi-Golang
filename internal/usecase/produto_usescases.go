package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type ProdutoUseCases struct {
	ProdutoRepository interfaces.ProdutoRepositoryInterface
}

func NewProdutoUseCases(ProdutoRepository interfaces.ProdutoRepositoryInterface) *ProdutoUseCases {
	return &ProdutoUseCases{ProdutoRepository: ProdutoRepository}
}

func (p *ProdutoUseCases) Inserir(produto *entity.Produto) error {
	return p.ProdutoRepository.Inserir(produto)
}

func (p *ProdutoUseCases) ObterTodosProdutos() (*[]entity.Produto, error) {
	return p.ProdutoRepository.ObterTodos()
}

func (p *ProdutoUseCases) ObterPorCategoria(categoria entity.CategoriaProduto) (*[]entity.Produto, error) {
	return p.ProdutoRepository.ObterPorCategoria(categoria)
}
