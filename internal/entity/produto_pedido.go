package entity

import "github.com/google/uuid"

type ProdutoPedido struct {
	PedidoId   uuid.UUID `gorm:"column:pedidoId"`
	ProdutoId  uuid.UUID `gorm:"column:produtoId"`
	Quantidade int       `gorm:"column:quantidade"`
}
