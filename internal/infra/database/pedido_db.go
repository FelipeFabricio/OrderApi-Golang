package database

import (
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
	err := p.Db.Find(&pedidos, "Status = ?", 3).Order("status DESC, data ASC").Error
	if err != nil {
		return nil, err
	}
	//p.Db.Raw("SELECT * FROM wonderfood.pedidos WHERE status <> 3 ORDER BY status DESC, data ASC;").Scan(&pedidos)
	return &pedidos, nil
}

// func (p *PedidoDb) Inserir(pedido *entity.Pedido) (*entity.Pedido, error) {

// 	p.Db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Create(pedido).Error; err != nil {
// 			return err
// 		}
// 		pedidoId := pedido.ID.String()
// 		var produtosPedido = []dto.ProdutoPedidoDto

// 		for index, produto := range pedido.Produtos {
// 			produtosPedido = []dto.ProdutoPedidoDto
// 			{
// 				{PedidoId: pedidoId, ProdutoId: pedido.Produtos[index].ID.String(), Quantidade: pedido.Produtos[index].Quantidade},

// 			}
// 		}

// 		if err != nil {
// 			return err
// 		}

// 		return pedido, nil
// 	})
// }
