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
	ClienteId    uuid.UUID       `json:"clienteId"`
	Valor        decimal.Decimal `json:"valor"`
	Status       StatusPedido    `json:"status"`
	Data         time.Time       `json:"data"`
	NumeroPedido uint64          `json:"numeroPedido"`
	Produtos     []Produto       `json:"produtos"`
}

func NewPedido(clienteId uuid.UUID, produtos *[]Produto) (*Pedido, error) {
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
	}
	pedido.CalcularValorPedido()

	return pedido, nil
}

func (p *Pedido) CalcularValorPedido() {
	var valor decimal.Decimal
	for _, produto := range p.Produtos {
		valor = valor.Add(produto.Valor)
	}
	p.Valor = valor
}
