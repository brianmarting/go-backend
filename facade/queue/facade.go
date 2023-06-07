package queue

import (
	"fmt"
	"go-backend/app/queue/rabbitmq"
	"os"

	"github.com/rs/zerolog/log"
)

func NewPublisher() *rabbitmq.Publisher {
	pub, err := rabbitmq.NewPublisher(getUrl())
	if err != nil {
		log.Panic().Msg("failed to create publisher")
	}

	return pub
}

func NewConsumer() *rabbitmq.Consumer {
	consumer, err := rabbitmq.NewConsumer(getUrl())
	if err != nil {
		log.Fatal().Msg("failed to create consumer")
	}

	return consumer
}

func getUrl() string {
	var (
		username = os.Getenv("RABBITMQ_USERNAME")
		password = os.Getenv("RABBITMQ_PASSWORD")
		host     = os.Getenv("RABBITMQ_HOST")
	)

	return fmt.Sprintf("amqp://%s:%s@%s:5672/", username, password, host)
}
