package publisher

import (
	"encoding/json"
	"gcnt/internal/stream"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

var Nats *natsPublisher

type INatsPublisher interface {
	Publish(subject string, payload interface{}) (err error)
	Request(subject string, payload []byte, structResult any) (err error)
}

type natsPublisher struct{}

func InitNatsPublisherInstance() {
	Nats = &natsPublisher{}
}

func (n *natsPublisher) Publish(subject string, payload interface{}) (err error) {
	data, _ := json.Marshal(payload)

	if stream.MessageBrokerInstance.Nats.IsConnected() {
		err = stream.MessageBrokerInstance.Nats.Publish(subject, data)
		if err != nil {
			log.Err(err).Send()
			return err
		}
		log.Info().Msg("Published message on subject " + subject)
	} else {
		err = nats.ErrConnectionClosed
		log.Err(err).Caller()
	}

	return
}

func (n *natsPublisher) Request(subject string, payload []byte, structResult any) (err error) {
	return
}
