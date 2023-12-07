package messaging

import (
	"log"

	amqp "github.com/streadway/amqp"
)

func OpenRabbitMQConnection(connectionString string) *amqp.Channel {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("Não foi possível se conectar ao RabbitMQ: %v", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("Não foi criar um Canal no RabbitMQ: %v", err)
	}

	_, err = channel.QueueDeclare(
		"pagamento-pedido",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Não foi possível declarar a fila: %v", err)
	}
	return channel
}
