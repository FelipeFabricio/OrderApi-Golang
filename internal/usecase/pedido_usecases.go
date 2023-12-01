package usecase

import (
	"fmt"
	"time"

	"github.com/felipefabricio/wonder-food/internal/dto"
	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
	"github.com/google/uuid"
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
			Status:       pedido.Status.ObterDescricaoStatusPedido(),
			Data:         pedido.Data,
			NumeroPedido: pedido.NumeroPedido,
		})
	}

	return &pedidosDto, nil
}

func (p *PedidoUseCases) Inserir(pedidoDto *dto.CriarPedidoInputDto) error {
	err := p.ValidarDadosPedido(pedidoDto)
	if err != nil {
		return err
	}

	pedido := entity.Pedido{
		ID:        uuid.New(),
		ClienteId: pedidoDto.ClienteId,
		Data:      time.Now(),
	}

	for _, produtoDto := range pedidoDto.Produtos {
		produtosPedido := entity.ProdutosPedido{
			PedidoId:   pedido.ID,
			ProdutoId:  produtoDto.ProdutoId,
			Quantidade: int(produtoDto.Quantidade),
		}
		pedido.Produtos = append(pedido.Produtos, produtosPedido)
	}

	p.CalcularValorFinalPedido(&pedido)
	return p.PedidoDb.Inserir(&pedido)
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

func (p *PedidoUseCases) ValidarDadosPedido(pedidoDto *dto.CriarPedidoInputDto) error {
	if len(pedidoDto.Produtos) == 0 {
		return entity.ErrPedidoSemProdutos
	}

	//TODO: Verificar se o Cliente existe
	if (pedidoDto.ClienteId == uuid.UUID{}) {
		return entity.ErrClienteIdInvalido
	}

	return nil
}
