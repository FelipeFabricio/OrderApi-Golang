package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Produto struct {
	ID        uuid.UUID        `json:"id"`
	Nome      string           `json:"nome"`
	Descricao string           `json:"descricao"`
	Categoria CategoriaProduto `json:"categoria"`
	Valor     decimal.Decimal  `json:"valor"`
}
