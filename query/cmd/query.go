package cmd

import (
	"fmt"
	"gcnt/config"
	"gcnt/internal/handler"
	"gcnt/internal/repository"
	"gcnt/internal/service"
	"gcnt/internal/stream"
	"gcnt/internal/stream/consumer"
	"gcnt/internal/stream/publisher"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(queryServiceCmd)
}

var queryServiceCmd = &cobra.Command{
	Use:   "query",
	Short: "Query Service",
	Long:  `Query Service`,
	Run: func(cmd *cobra.Command, args []string) {
		BootStrap()
	},
}

func BootStrap() {
	//Init Storage
	db := &repository.Db{}
	db.InitDatabaseInstance("typesense")

	//Init Service
	service.InitArticleServiceInstance()

	//Init Repo
	repository.InitArticleRepositoryInstance()

	//Init message broker
	err := stream.InitMessageBrokerInstance()
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
		return
	}

	//Init publisher
	publisher.InitNatsPublisherInstance()

	//Init Consumer
	consumer.InitNATSConsumer()

	//Init & Start HTTP server
	app := handler.SetupAPIRouter()
	err = app.Listen(fmt.Sprintf(":%d", config.Instance.CommandServicePort))
	if err != nil {
		log.Panic().Caller().Err(err)
	}
}
