package dto

import (
	"time"

	"github.com/felipefabricio/wonder-food/internal/entity"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CriarPedidoInputDto struct {
	ClienteId uuid.UUID          `json:"clienteId"`
	Produto   []ProdutoPedidoDto `json:"produtos"`
}

type CriarPedidoOutputDto struct {
	ID           uuid.UUID           `json:"id"`
	ClienteId    uuid.UUID           `json:"clienteId"`
	Valor        decimal.Decimal     `json:"valor"`
	Status       entity.StatusPedido `json:"status"`
	Data         time.Time           `json:"data"`
	NumeroPedido uint64              `json:"numeroPedido"`
}

type ObterPedidosOutputDto struct {
	ID           uuid.UUID           `json:"id"`
	ClienteId    uuid.UUID           `json:"clienteId"`
	Valor        decimal.Decimal     `json:"valor"`
	Status       entity.StatusPedido `json:"status"`
	Data         time.Time           `json:"data"`
	NumeroPedido uint64              `json:"numeroPedido"`
}

type ProdutoPedidoDto struct {
	PedidoId   string `json:"pedidoId"`
	ProdutoId  string `json:"produtoId"`
	Quantidade uint8  `json:"quantidade"`
}
