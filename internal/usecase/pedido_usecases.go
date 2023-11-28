package usecase

import (
	"fmt"

	"github.com/felipefabricio/wonder-food/internal/dto"
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
	"github.com/shopspring/decimal"
)

type PedidoUseCases struct {
	PedidoDb  interfaces.PedidoDbInterface
	ProdutoDb interfaces.ProdutoDbInterface
}

func NewPedidoUseCases(pedidoDb interfaces.PedidoDbInterface, produtoDb interfaces.ProdutoDbInterface) *PedidoUseCases {
	return &PedidoUseCases{
		PedidoDb:  pedidoDb,
		ProdutoDb: produtoDb,
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

func (p *PedidoUseCases) Inserir(pedidoDto *dto.CriarPedidoInputDto) error {
	var pedido *entity.Pedido

	if len(pedidoDto.Produtos) == 0 {
		return entity.ErrPedidoSemProdutos
	}

	for _, produtoDto := range pedidoDto.Produtos {
		produtosPedido := entity.ProdutoPedido{
			PedidoId:   pedido.ID,
			ProdutoId:  produtoDto.ProdutoId,
			Quantidade: int(produtoDto.Quantidade),
		}

		pedido.Produtos = append(pedido.Produtos, produtosPedido)
	}

	pedidoValidado, err := pedido.NewPedido(pedido.ClienteId, &pedido.Produtos)
	if err != nil {
		return err
	}
	p.CalcularValorFinalPedido(pedidoValidado)
	return p.PedidoDb.Inserir(pedidoValidado)
}

func (p *PedidoUseCases) CalcularValorFinalPedido(pedido *entity.Pedido) {
	var valorFinal decimal.Decimal

	for _, produtosPedido := range pedido.Produtos {
		produto, err := p.ProdutoDb.ObterPorId(produtosPedido.ProdutoId.String())
		if err != nil {
			fmt.Println("Não existe um Produto com o Id informado.", err)
			return
		}

		valorFinal = valorFinal.Add(produto.Valor.Mul(decimal.NewFromInt(int64(produtosPedido.Quantidade))))
	}
	pedido.Valor = valorFinal
}
