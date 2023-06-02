package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Connection struct {
	*amqp.Channel
}

func GetConnection(url string) (Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return Connection{}, err
	}

	channel, err := conn.Channel()
	return Connection{
		Channel: channel,
	}, err
}
