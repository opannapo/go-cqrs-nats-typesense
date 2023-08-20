package streams

import (
	"gcnt/config"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

var MessageBrokerInstance *messageBroker

type messageBroker struct {
	Nats *nats.Conn
}

func InitMessageBrokerInstance() (err error) {
	nats, err := natsClient()
	if err != nil {
		log.Err(err).Send()
		return
	}

	MessageBrokerInstance = &messageBroker{Nats: nats}
	return
}

func natsClient() (clientConn *nats.Conn, err error) {
	clientConn, err = nats.Connect(config.Instance.NatsAddress)
	if err != nil {
		log.Err(err).Send()
		return
	}

	return
}
