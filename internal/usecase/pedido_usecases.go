package usecase

import (
	"github.com/felipefabricio/wonder-food/internal/dto"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
)

type PedidoUseCases struct {
	PedidoDb interfaces.PedidoDbInterface
}

func NewPedidoUseCases(pedidoDb interfaces.PedidoDbInterface) *PedidoUseCases {
	return &PedidoUseCases{
		PedidoDb: pedidoDb,
	}
}

func (p *PedidoUseCases) ObterPedidosEmAberto() (*[]dto.ObterPedidosOutputDto, error) {
	pedidos, err := p.PedidoDb.ObterPedidosEmAberto()
	if err != nil {
		return nil, err
	}

	var pedidosDto []dto.ObterPedidosOutputDto
	for _, pedido := range *pedidos {
		pedidosDto = append(pedidosDto, dto.ObterPedidosOutputDto{
			ID:           pedido.ID,
			ClienteId:    pedido.ClienteId,
			Valor:        pedido.Valor,
			Status:       pedido.Status,
			Data:         pedido.Data,
			NumeroPedido: pedido.NumeroPedido,
		})
	}

	return &pedidosDto, nil
}

// func (p *PedidoUseCases) Inserir(pedidoDto *dto.CriarPedidoInputDto, produtos *[]dto.ProdutoDto) (*dto.CriarPedidoOutputDto, error) {
// 	var pedido *entity.Pedido
// 	pedido = pedido.NewPedido(pedidoDto.ClienteId, produtos)
// 	return p.PedidoDb.Inserir(pedido, produtos)
// }
