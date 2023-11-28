package database

import (
	"errors"

	"github.com/felipefabricio/wonder-food/internal/entity"
	"gorm.io/gorm"
)

type PedidoDb struct {
	Db *gorm.DB
}

func NewPedidoDb(db *gorm.DB) *PedidoDb {
	return &PedidoDb{
		Db: db,
	}
}

func (p *PedidoDb) ObterPedidosEmAberto() (*[]entity.Pedido, error) {
	var pedidos []entity.Pedido
	p.Db.Raw("SELECT id, clienteId, valor, status, data, numeroPedido FROM wonderfood.pedidos WHERE status <> 3 ORDER BY status DESC, data ASC;").Scan(&pedidos)
	return &pedidos, nil
}

func (p *PedidoDb) Inserir(pedido *entity.Pedido) error {
	tx := p.Db.Begin()

	if err := tx.Create(pedido).Error; err != nil {
		tx.Rollback()
		return errors.New("Erro ao inserir Pedido")
	}

	for i := range pedido.Produtos {
		pedido.Produtos[i].PedidoId = pedido.ID
		if err := tx.Create(&pedido.Produtos[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
