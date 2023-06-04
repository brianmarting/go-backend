package facade

import (
	"github.com/rs/zerolog/log"
	"go-backend/interfaces/queue"
	"go-backend/queue/rabbitmq"
	"sync"
)

var once sync.Once

var publisher rabbitmq.Publisher

func init() {
	once.Do(func() {
		publisher = *newPublisher("amqp://guest:guest@localhost:5672/")
	})
}

func GetPublisher() queue.Publisher {
	return publisher
}

func newPublisher(url string) *rabbitmq.Publisher {
	pub, err := rabbitmq.NewPublisher(url)
	if err != nil {
		log.Panic().Msg("failed to create publisher")
	}

	return pub
}

func NewConsumer(url string) *rabbitmq.Consumer {
	consumer, err := rabbitmq.NewConsumer(url)
	if err != nil {
		log.Fatal().Msg("failed to create consumer")
	}

	return consumer
}
