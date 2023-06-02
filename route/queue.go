package route

import (
	"github.com/rs/zerolog/log"
	"go-backend/queue/rabbitmq"
)

func NewPublisher(url string) *rabbitmq.Publisher {
	pub, err := rabbitmq.NewPublisher(url)
	if err != nil {
		log.Panic().Msg("Failed to create publisher")
	}

	return pub
}
