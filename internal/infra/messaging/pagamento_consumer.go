package messaging

import (
	"log"

	amqp "github.com/streadway/amqp"
)

func Consume(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"pagamento-pedido",
		"wonderfood-consumer",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalf("Não foi possível consumir a fila de pagamento-pedido: %v", err)
	}

	for msg := range msgs {
		log.Printf("Mensagem recebida: %s", msg.Body)
		msg.Ack(false)
	}
}
