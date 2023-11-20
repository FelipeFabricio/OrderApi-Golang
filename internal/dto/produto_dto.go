package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProdutoDto struct {
	ID        uuid.UUID       `json:"id"`
	Nome      string          `json:"nome"`
	Descricao string          `json:"descricao"`
	Categoria uint8           `json:"categoria"`
	Valor     decimal.Decimal `json:"valor"`
}
