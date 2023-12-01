package entity

import "github.com/google/uuid"

type ProdutosPedido struct {
	PedidoId   uuid.UUID `gorm:"column:pedidoId"`
	ProdutoId  uuid.UUID `gorm:"column:produtoId"`
	Quantidade int       `gorm:"column:quantidade"`
}

func (ProdutosPedido) TableName() string {
	return "produtos_pedido"
}
