package queue

import (
	"fmt"
	rabbitmq2 "go-backend/internal/app/queue/rabbitmq"
	"os"

	"github.com/rs/zerolog/log"
)

func NewPublisher() *rabbitmq2.Publisher {
	pub, err := rabbitmq2.NewPublisher(getUrl())
	if err != nil {
		log.Panic().Msg("failed to create publisher")
	}

	return pub
}

func NewConsumer() *rabbitmq2.Consumer {
	consumer, err := rabbitmq2.NewConsumer(getUrl())
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
