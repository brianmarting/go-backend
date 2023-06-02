package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Message struct {
	*amqp.Delivery
}

func (m Message) GetBytes() []byte {
	return m.Delivery.Body
}
