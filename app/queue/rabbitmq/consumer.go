package rabbitmq

import (
	"go-backend/app/queue"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	*Connection
}

func NewConsumer(url string) (*Consumer, error) {
	conn, err := GetConnection(url)
	if err != nil {
		return &Consumer{}, nil
	}

	return &Consumer{
		Connection: &conn,
	}, nil
}

func (c Consumer) StartConsuming(queueName, key string) (<-chan queue.Message, error) {
	_, err := c.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = c.channel.QueueBind(
		queueName,
		key,
		"amq.direct",
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	amqpDeliveries, err := c.channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	messageChannel := make(chan queue.Message)

	go func() {
		defer close(messageChannel)
		for delivery := range amqpDeliveries {
			message := convertToMessage(delivery)
			messageChannel <- message
		}
	}()

	return messageChannel, nil
}

func convertToMessage(delivery amqp.Delivery) Message {
	return Message{
		Delivery: &delivery,
	}
}
