package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	*Connection
}

func (p Publisher) Publish(ctx context.Context, routingKey string, data []byte) error {
	return p.Channel.PublishWithContext(
		ctx,
		"amq.direct",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
			Timestamp:   time.Now(),
		},
	)
}

func NewPublisher(url string) (*Publisher, error) {
	conn, err := GetConnection(url)
	if err != nil {
		return &Publisher{}, nil
	}

	return &Publisher{
		Connection: &conn,
	}, nil
}
