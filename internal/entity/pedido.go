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
	ID           uuid.UUID       `json:"id"`
	ClienteId    uuid.UUID       `gorm:"column:clienteId"`
	Valor        decimal.Decimal `json:"valor"`
	Status       StatusPedido    `json:"status"`
	Data         time.Time       `json:"data"`
	NumeroPedido int             `gorm:"column:numeroPedido"`
	Produtos     []ProdutoPedido `gorm:"foreignKey:PedidoId;references:ID"`
}

func (p *Pedido) NewPedido(clienteId uuid.UUID, produtos *[]ProdutoPedido) (*Pedido, error) {
	if len(*produtos) == 0 {
		return nil, ErrPedidoSemProdutos
	}

	if clienteId == uuid.Nil {
		return nil, ErrClienteIdInvalido
	}

	pedido := &Pedido{
		ID:        uuid.New(),
		ClienteId: clienteId,
		Produtos:  *produtos,
		Status:    0,
	}

	return pedido, nil
}
