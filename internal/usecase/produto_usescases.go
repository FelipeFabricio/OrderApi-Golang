package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/dto"
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type ProdutoUseCases struct {
	ProdutoDb interfaces.ProdutoDbInterface
}

func NewProdutoUseCases(produtoDb interfaces.ProdutoDbInterface) *ProdutoUseCases {
	return &ProdutoUseCases{ProdutoDb: produtoDb}
}

func (p *ProdutoUseCases) Inserir(produto *dto.CriarProdutoInputDto) error {
	var novoProduto entity.Produto
	produtoValidado, err := novoProduto.NewProduto(produto.Nome, produto.Descricao, produto.Categoria, produto.Valor)
	if err != nil {
		return err
	}
	return p.ProdutoDb.Inserir(produtoValidado)
}

func (p *ProdutoUseCases) ObterTodosProdutos() (*[]dto.ObterProdutoOutputDto, error) {
	produtos, err := p.ProdutoDb.ObterTodos()
	if err != nil {
		return nil, err
	}

	var produtosOutput []dto.ObterProdutoOutputDto
	for _, produto := range *produtos {
		produtosOutput = append(produtosOutput, dto.ObterProdutoOutputDto{
			ID:        produto.ID,
			Nome:      produto.Nome,
			Descricao: produto.Descricao,
			Categoria: produto.Categoria,
			Valor:     produto.Valor,
		})
	}
	return &produtosOutput, nil
}

func (p *ProdutoUseCases) ObterPorCategoria(categoria entity.CategoriaProduto) (*[]dto.ObterProdutoOutputDto, error) {
	produtos, err := p.ProdutoDb.ObterPorCategoria(categoria)
	if err != nil {
		return nil, err
	}

	var produtosOutput []dto.ObterProdutoOutputDto
	for _, produto := range *produtos {
		produtosOutput = append(produtosOutput, dto.ObterProdutoOutputDto{
			ID:        produto.ID,
			Nome:      produto.Nome,
			Descricao: produto.Descricao,
			Categoria: produto.Categoria,
			Valor:     produto.Valor,
		})
	}
	return &produtosOutput, nil
}
