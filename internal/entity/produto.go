package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TODO: Criar validações detalhadas como: tamanho do campo, required, etc
var (
	ErrCategoriaProdutoInvalida = errors.New("Categoria de produto inválida!")
)

type Produto struct {
	ID        uuid.UUID        `json:"id"`
	Nome      string           `json:"nome"`
	Descricao string           `json:"descricao"`
	Categoria CategoriaProduto `json:"categoria"`
	Valor     decimal.Decimal  `json:"valor"`
}

func (p *Produto) NewProduto(nome, descricao string, categoria CategoriaProduto, valor decimal.Decimal) (*Produto, error) {

	if !p.ValidarCategoriaProduto() {
		return nil, ErrCategoriaProdutoInvalida
	}

	novoProduto := &Produto{
		ID:        uuid.New(),
		Nome:      nome,
		Descricao: descricao,
		Categoria: categoria,
		Valor:     valor,
	}
	return novoProduto, nil
}

func (p *Produto) ValidarCategoriaProduto() bool {
	return p.Categoria >= Lanche && p.Categoria <= Acompanhamento
}
