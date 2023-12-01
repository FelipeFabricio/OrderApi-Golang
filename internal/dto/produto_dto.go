package dto

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CriarProdutoInputDto struct {
	Nome      string                  `json:"nome"`
	Descricao string                  `json:"descricao"`
	Categoria entity.CategoriaProduto `json:"categoria"`
	Valor     decimal.Decimal         `json:"valor"`
}

type ObterProdutoOutputDto struct {
	ID        uuid.UUID       `json:"id"`
	Nome      string          `json:"nome"`
	Descricao string          `json:"descricao"`
	Categoria string          `json:"categoria"`
	Valor     decimal.Decimal `json:"valor"`
}
