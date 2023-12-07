package database

import (
	"errors"

	"github.com/felipefabricio/wonder-food/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrPedidoNaoEncontrado = errors.New("pedido não encontrado")
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

func (p *PedidoDb) ObterPorNumeroPedido(numeroPedido int) (*entity.Pedido, error) {
	var pedido entity.Pedido
	p.Db.Where("numeroPedido = ?", numeroPedido).First(&pedido)

	if pedido.ID == uuid.Nil {
		return nil, ErrPedidoNaoEncontrado
	}
	return &pedido, nil
}

func (p *PedidoDb) AtualizarStatusPagamento(numeroPedido int, status entity.StatusPedido) error {
	pedido, err := p.ObterPorNumeroPedido(numeroPedido)
	if err != nil {
		return ErrPedidoNaoEncontrado
	}
	pedido.Status = status
	p.Db.Save(&pedido)
	return nil
}

func (p *PedidoDb) DeletarPedido(numeroPedido int) error {
	_, err := p.ObterPorNumeroPedido(numeroPedido)
	if err != nil {
		return ErrPedidoNaoEncontrado
	}
	p.Db.Exec("DELETE FROM wonderfood.pedidos WHERE numeroPedido = ?", numeroPedido)
	return nil
}
