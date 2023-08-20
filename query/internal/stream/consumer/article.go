package consumer

import (
	"context"
	"encoding/json"
	"gcnt/internal/schema"
	"gcnt/internal/service"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func articleCreated(msg *nats.Msg) {
	payload := schema.MessageConsume{}
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		log.Err(err).Caller().Send()
	}
	log.Info().Caller().Msgf("Message in : %+v", payload)

	ctx := context.Background()

	err := service.ArticleServiceInstance.Upsert(ctx, payload)
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}
}
