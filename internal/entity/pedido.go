package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var (
	ErrClienteIdInvalido = errors.New("Não foi encontrado um Cliente com  o Id informado")
	ErrPedidoSemProdutos = errors.New("Não é possível inserir um Pedido sem Produtos")
)

type Pedido struct {
	ID           uuid.UUID        `json:"id"`
	ClienteId    uuid.UUID        `gorm:"column:clienteId"`
	Valor        decimal.Decimal  `json:"valor"`
	Status       StatusPedido     `json:"status"`
	Data         time.Time        `json:"data"`
	NumeroPedido int              `gorm:"column:numeroPedido"`
	Produtos     []ProdutosPedido `gorm:"foreignKey:PedidoId;references:ID"`
}
