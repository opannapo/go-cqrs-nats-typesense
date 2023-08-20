package consumer

import (
	"encoding/json"
	"gcnt/internal/schema"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func articleCreated(msg *nats.Msg) {
	payload := schema.MessageConsume{}
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		log.Err(err).Caller().Send()
	}

	log.Printf("Message in : %+v", payload)
}
