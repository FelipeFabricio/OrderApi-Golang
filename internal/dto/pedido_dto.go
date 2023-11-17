package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PedidoDto struct {
	ID        uuid.UUID       `json:"id"`
	ClienteId uuid.UUID       `json:"clienteId"`
	Valor     decimal.Decimal `json:"valor"`
	Status    uint8           `json:"status"`
	Data      time.Time       `json:"data"`
}
