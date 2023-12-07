package messaging

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
	amqp "github.com/streadway/amqp"
)

type RetornoPagamentoPedido struct {
	NumeroPedido       int  `json:"numeroPedido"`
	PagamentoRealizado bool `json:"pagamentoRealizado"`
}

type PagamentoConsumer struct {
	PedidoDb interfaces.PedidoDbInterface
}

func NewPagamentoConsumer(pedidoDbb interfaces.PedidoDbInterface) *PagamentoConsumer {
	return &PagamentoConsumer{
		PedidoDb: pedidoDbb,
	}
}

func (c *PagamentoConsumer) Consume(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"pagamento-pedido",
		"wonderfood-consumer",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalf("Não foi possível consumir a fila 'pagamento-pedido': %v", err)
	}

	var retornoPagamento RetornoPagamentoPedido
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &retornoPagamento)
		if err != nil {
			fmt.Printf("Não foi possível deserializar a mensagem: %v", err)
		}

		if retornoPagamento.PagamentoRealizado {
			c.PedidoDb.AtualizarStatusPagamento(retornoPagamento.NumeroPedido, entity.Recebido)
		} else {
			c.PedidoDb.DeletarPedido(retornoPagamento.NumeroPedido)
		}

		log.Printf("Retorno do Pagamento do pedido %d recebido! Pagamento efetuado: %v", retornoPagamento.NumeroPedido, retornoPagamento.PagamentoRealizado)
		msg.Ack(false)
	}
}
