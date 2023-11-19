package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Pedido struct {
	ID           uuid.UUID       `json:"id"`
	ClienteId    uuid.UUID       `json:"clienteId"`
	Valor        decimal.Decimal `json:"valor"`
	Status       StatusPedido    `json:"status"`
	Data         time.Time       `json:"data"`
	NumeroPedido uint64          `json:"numeroPedido"`
}
