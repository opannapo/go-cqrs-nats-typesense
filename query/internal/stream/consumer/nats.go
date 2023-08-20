package consumer

import (
	"gcnt/internal/stream"
	"github.com/rs/zerolog/log"
)

func InitNATSConsumer() {
	mb := stream.MessageBrokerInstance

	_, err := mb.Nats.Subscribe("article.created", articleCreated)
	if err != nil {
		log.Err(err).Send()
	}
}
